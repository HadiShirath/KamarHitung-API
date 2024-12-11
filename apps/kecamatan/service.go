package kecamatan

import (
	"context"
	"fmt"
	"kamar-hitung/apps/auth"
	"kamar-hitung/infra/response"
	"kamar-hitung/internal/config"
	"strconv"

	"github.com/google/uuid"
)

type Repository interface {
	GetListKecamatanCode(ctx context.Context) (kecamatans []Kecamatan, err error)
	GetVoterKecamatan(ctx context.Context, codeKecamatan string) (kecamatan Kecamatan, err error)
	GetAllVoter(ctx context.Context) (kecamatan Kecamatan, err error)
	GetListKecamatan(ctx context.Context) (kecamatans []Kecamatan, err error)
	GetListKelurahanFromKecamatan(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error)
	deleteDataKecamatanKelurahanTPS(ctx context.Context) (err error)
	createKecamatan(ctx context.Context, model Output) (err error)
	createKelurahan(ctx context.Context, model KelurahanOutput) (err error)
	createTps(ctx context.Context, model TpsOutput) error
}

type service struct {
	repo     Repository
	repoAuth auth.Repository
}

func NewService(repo Repository, repoAuth auth.Repository) service {
	return service{
		repo:     repo,
		repoAuth: repoAuth,
	}
}

func (s service) GetVoterKecamatan(ctx context.Context, codeKecamatan string) (kecamatan Kecamatan, err error) {

	model, err := s.repo.GetVoterKecamatan(ctx, codeKecamatan)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) AllVoter(ctx context.Context) (kecamatan Kecamatan, err error) {

	model, err := s.repo.GetAllVoter(ctx)
	if err != nil {
		return
	}

	return model, nil
}

func (s service) GetListKecamatan(ctx context.Context) (kecamatans []Kecamatan, err error) {

	kecamatans, err = s.repo.GetListKecamatan(ctx)
	if err != nil {
		if err == response.ErrNotFound {
			return []Kecamatan{}, nil
		}
		return
	}

	if len(kecamatans) == 0 {
		return []Kecamatan{}, nil
	}
	return
}

func (s service) GetListKelurahanFromKecamatan(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error) {

	kelurahans, err = s.repo.GetListKelurahanFromKecamatan(ctx, codeKecamatan)

	if err != nil {
		if err == response.ErrNotFound {
			return []Kelurahan{}, nil
		}
		return
	}

	if len(kelurahans) == 0 {
		return []Kelurahan{}, nil
	}
	return
}

func (s service) GetListKecamatanCode(ctx context.Context) (kecamatans []Kecamatan, err error) {

	kecamatans, err = s.repo.GetListKecamatanCode(ctx)

	if err != nil {
		if err == response.ErrNotFound {
			return []Kecamatan{}, nil
		}
		return
	}

	if len(kecamatans) == 0 {
		return []Kecamatan{}, nil
	}
	return
}

func (s service) ImportCSV(ctx context.Context, records [][]string) (err error) {

	// Sample data, jika CSV header diabaikan
	dataArray := make([]Data, 0)

	err = validateFormatCSV(records[0])
	if err != nil {
		return
	}

	for _, record := range records[1:] { // Mengabaikan header
		totalVoters, _ := strconv.Atoi(record[3]) // Ubah kolom ke-4 ke integer
		dataArray = append(dataArray, Data{
			KecamatanName: record[0],
			KelurahanName: record[1],
			TpsName:       record[2],
			TotalVoters:   totalVoters,
			Fullname:      record[4],
			Username:      record[5],
		})

	}

	// Maps to store IDs and codes
	kecamatanDataMap := make(map[string]*Output)
	kelurahanDataMap := make(map[string]*KelurahanOutput)
	tpsDataMap := make([]*TpsOutput, 0)

	// Proses data
	err = s.repo.deleteDataKecamatanKelurahanTPS(context.Background())
	if err != nil {
		return
	}

	for _, data := range dataArray {
		// Cek jika Kecamatan sudah ada
		if _, exists := kecamatanDataMap[data.KecamatanName]; !exists {
			kecamatanDataMap[data.KecamatanName] = &Output{
				KecamatanID:          generateUUID(),
				KecamatanName:        data.KecamatanName,
				TotalVotersKecamatan: 0,
				TotalTpsKecamatan:    0,
				CodeKecamatan:        generateCode(),
			}
		}

		kecamatanEntry := kecamatanDataMap[data.KecamatanName]

		// Hitung total pemilih dan TPS untuk Kecamatan
		kecamatanEntry.TotalVotersKecamatan += data.TotalVoters
		kecamatanEntry.TotalTpsKecamatan++

		// Cek jika Kelurahan sudah ada
		kelurahanKey := fmt.Sprintf("%s-%s", kecamatanEntry.KecamatanID, data.KelurahanName)
		if _, exists := kelurahanDataMap[kelurahanKey]; !exists {
			kelurahanDataMap[kelurahanKey] = &KelurahanOutput{
				KelurahanID:          generateUUID(),
				KecamatanID:          kecamatanEntry.KecamatanID,
				KelurahanName:        data.KelurahanName,
				TotalVotersKelurahan: 0,
				TotalTpsKelurahan:    0,
				CodeKelurahan:        generateCode(),
			}
		}

		kelurahanEntry := kelurahanDataMap[kelurahanKey]

		// Hitung total pemilih untuk Kelurahan
		kelurahanEntry.TotalVotersKelurahan += data.TotalVoters
		kelurahanEntry.TotalTpsKelurahan++ // Hitung TPS per Kelurahan

		// Hapus Semua data Kecamatan, Kelurahan, TPS & Auth

		// Daftar User
		reqUser := auth.RegisterRequestPayload{
			Fullname: data.Fullname,
			Username: data.Username,
			Password: generatePassword(),
		}

		authEntity := auth.NewFromRegisterRequest(reqUser)
		if err = authEntity.Validate(); err != nil {
			return
		}

		if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
			return
		}

		err = s.repoAuth.CreateAuth(ctx, authEntity)
		if err != nil {
			return
		}

		// Buat entri TPS
		tpsEntry := &TpsOutput{
			TpsID:         generateUUID(),
			KecamatanID:   kecamatanEntry.KecamatanID,
			KelurahanID:   kelurahanEntry.KelurahanID,
			TpsName:       data.TpsName,
			TotalVoters:   data.TotalVoters,
			Paslon1:       0,  // Set sesuai data jika ada
			Paslon2:       0,  // Set sesuai data jika ada
			Paslon3:       0,  // Set sesuai data jika ada
			Paslon4:       0,  // Set sesuai data jika ada
			SuaraSah:      0,  // Set sesuai data jika ada
			SuaraTidakSah: 0,  // Set sesuai data jika ada
			Photo:         "", // Set sesuai kebutuhan
			Code:          generateCode(),
			UserID:        authEntity.PublicId.String(),
		}
		tpsDataMap = append(tpsDataMap, tpsEntry)
	}

	// Sisipkan data ke dalam tabel kecamatan
	for _, model := range kecamatanDataMap {
		err := s.repo.createKecamatan(context.Background(), *model)
		if err != nil {
			fmt.Println("Error inserting data:", err)
			continue // skip to next iteration
		}
	}

	// Sisipkan data ke dalam tabel kelurahan
	for _, model := range kelurahanDataMap {
		err := s.repo.createKelurahan(context.Background(), *model)
		if err != nil {
			fmt.Println("Error inserting kelurahan data:", err)
		}
	}

	// Sisipkan data ke dalam tabel tps
	for _, model := range tpsDataMap {
		err := s.repo.createTps(context.Background(), *model)
		if err != nil {
			fmt.Println("Error inserting tps data:", err)
		}
	}

	return nil

}

// Fungsi untuk menghasilkan UUID
func generateUUID() string {
	return uuid.New().String()
}

// Fungsi untuk menghasilkan kode acak 6 karakter
func generateCode() string {
	code := uuid.New().String()[:6]
	return code
}

func generatePassword() string {
	code := uuid.New().String()[:4]
	return code
}
