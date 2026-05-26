package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/aquasecurity/table"
)

type ServiceAccount struct {
	NamaLayanan       string
	NamaPengguna      string
	AlamatEmail       string
	KataSandi         string
	KekuatanKataSandi string
	TanggalPembaruan  time.Time
	WaktuInput        time.Time
}

var ServiceAccounts [100]ServiceAccount

func PasswordStrenth(password string) string {
	var hasLetter, hasNumber, hasSpecial bool
	var char rune

	for _, char = range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasNumber = true
		default:
			hasSpecial = true
		}
	}

	switch {
	case hasLetter && hasNumber && hasSpecial:
		return "Kuat"
	case hasLetter && hasNumber:
		return "Sedang"
	default:
		return "Lemah"
	}
}

// INPUTING FUNCTION
func InputText(prompt string) string {
	var scanner = bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// HANDLING FUNCTIONS
func HandleError(err error) bool {
	if err != nil {
		fmt.Println("\nERROR:", err)
		return true
	}

	return false
}

func HandleIndex(indexInput string, serviceAccountsCount *int) (int, bool) {
	var serviceAccountIndex int
	var err error

	serviceAccountIndex, err = strconv.Atoi(indexInput)

	if err != nil || serviceAccountIndex < 1 || serviceAccountIndex > *serviceAccountsCount {
		HandleError(errors.New("Nomor akun tidak valid."))
		return 0, false
	}

	return serviceAccountIndex, true
}

// CRUDING FUNCTIONS
func AddServiceAccount(serviceName, userName, emailAddress, password string, serviceAccountsCount *int) {
	var serviceAccount ServiceAccount

	serviceAccount.NamaLayanan = serviceName
	serviceAccount.NamaPengguna = userName
	serviceAccount.AlamatEmail = emailAddress
	serviceAccount.KataSandi = password
	serviceAccount.KekuatanKataSandi = PasswordStrenth(password)
	serviceAccount.TanggalPembaruan = time.Now()
	serviceAccount.WaktuInput = time.Now()

	ServiceAccounts[*serviceAccountsCount] = serviceAccount
	*serviceAccountsCount++

	fmt.Println("\nINFO: Akun layanan berhasil ditambahkan.")
}

func EditServiceAccount(serviceAccountIndex int, serviceName, userName, emailAddress, password string) {
	var currentIndex = serviceAccountIndex - 1

	ServiceAccounts[currentIndex].NamaLayanan = serviceName
	ServiceAccounts[currentIndex].NamaPengguna = userName
	ServiceAccounts[currentIndex].AlamatEmail = emailAddress
	ServiceAccounts[currentIndex].KataSandi = password
	ServiceAccounts[currentIndex].KekuatanKataSandi = PasswordStrenth(password)
	ServiceAccounts[currentIndex].TanggalPembaruan = time.Now()

	fmt.Println("\nINFO: Akun layanan berhasil diedit.")
}

func DeleteServiceAccount(serviceAccountIndex int, serviceAccountsCount *int) {
	var i int

	for i = serviceAccountIndex - 1; i < *serviceAccountsCount; i++ {
		ServiceAccounts[i] = ServiceAccounts[i+1]
	}

	ServiceAccounts[*serviceAccountsCount-1] = ServiceAccount{}
	*serviceAccountsCount--

	fmt.Println("\nINFO: Akun layanan berhasil dihapus.")
}

func ShowServiceAccounts(serviceName, userName, emailAddress, password string, serviceAccountsCount *int) {
	var i int
	var t *table.Table

	t = table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Nama Layanan", "Nama Pengguna", "Email", "Kata Sandi", "Kekuatan Kata Sandi", "Dibuat Pada", "Diperbarui Pada")

	for i = 0; i < *serviceAccountsCount; i++ {
		t.AddRow(
			strconv.Itoa(i+1),
			ServiceAccounts[i].NamaLayanan,
			ServiceAccounts[i].NamaPengguna,
			ServiceAccounts[i].AlamatEmail,
			ServiceAccounts[i].KataSandi,
			ServiceAccounts[i].KekuatanKataSandi,
			ServiceAccounts[i].WaktuInput.Format(time.DateTime),
			ServiceAccounts[i].TanggalPembaruan.Format(time.DateTime),
		)
	}

	fmt.Println()
	t.Render()
}

// SEARCHING FUNCTIONS
func SeqSearchByServiceName(keyword string, serviceAccountsCount *int) {
	var i, found int
	var t *table.Table

	keyword = strings.ToLower(keyword)

	t = table.New(os.Stdout)
	t.SetRowLines(false)
	t.SetHeaders("#", "Nama Layanan", "Nama Pengguna", "Email", "Kata Sandi", "Kekuatan Kata Sandi", "Dibuat Pada", "Diperbarui Pada")

	for i = 0; i < *serviceAccountsCount; i++ {
		if strings.Contains(strings.ToLower(ServiceAccounts[i].NamaLayanan), keyword) {
			t.AddRow(
				strconv.Itoa(i+1),
				ServiceAccounts[i].NamaLayanan,
				ServiceAccounts[i].NamaPengguna,
				ServiceAccounts[i].AlamatEmail,
				ServiceAccounts[i].KataSandi,
				ServiceAccounts[i].KekuatanKataSandi,
				ServiceAccounts[i].WaktuInput.Format(time.DateTime),
				ServiceAccounts[i].TanggalPembaruan.Format(time.DateTime),
			)

			found++
		}
	}

	if found == 0 {
		HandleError(errors.New("Nama layanan tidak ditemukan."))
	} else {
		fmt.Println("\n\n=== HASIL PENCARIAN ===")
		fmt.Println()
		t.Render()
	}
}

func BinSearchByServiceName(keyword string, serviceAccountsCount *int) {}

// SORTING FUNCTIONS
func SortByServiceName() {}

func SortByInputTime() {}

// TODO: Simpan dan load data akun dari file json

func main() {
	var p, se, serviceAccountsCount, serviceAccountIndex int
	var pInput, indexInput, seInput, serviceName, userName, emailAddress, password, keyword string
	var ok bool
	var err error

	fmt.Println("\n=== SECUREPASS - APLIKASI PENGELOLA KATA SANDI PRIBADI ===")

	for {
		fmt.Println("\n\nMenu utama:")
		fmt.Println("\n1. Lihat semua akun layanan")
		fmt.Println("2. Tambah akun layanan")
		fmt.Println("3. Edit akun layanan")
		fmt.Println("4. Hapus akun layanan")
		fmt.Println("5. Cari nama layanan")
		fmt.Println("6. Urutkan akun layanan")
		fmt.Println("0. Keluar")

		pInput = InputText("\n? Pilih menu\n> ")
		p, err = strconv.Atoi(pInput)
		for err != nil {
			HandleError(errors.New("Pilihan harus berupa angka."))

			pInput = InputText("\n? Pilih menu\n> ")
			p, err = strconv.Atoi(pInput)
		}
		fmt.Println()

		switch p {
		case 1:
			fmt.Println("\n=== DAFTAR SEMUA AKUN LAYANAN ===")

			if serviceAccountsCount == 0 {
				fmt.Println("\nINFO: Belum ada akun layanan.")
			} else {
				ShowServiceAccounts(serviceName, userName, emailAddress, password, &serviceAccountsCount)

				// TODO: Menampilkan statistik klasifikasi tingkat keamanan kata sandi
				fmt.Println("\nJumlah akun layanan yang tersimpan :", serviceAccountsCount)
			}
		case 2:
			fmt.Println("\n=== TAMBAH AKUN LAYANAN ===")

			serviceName = InputText("\n? Nama layanan\n> ")
			for serviceName == "" {
				HandleError(errors.New("Nama layanan tidak boleh kosong."))
				serviceName = InputText("\n? Nama layanan\n> ")
			}

			userName = InputText("\n? Nama pengguna\n> ")
			for userName == "" {
				HandleError(errors.New("Nama pengguna tidak boleh kosong."))
				userName = InputText("\n? Nama pengguna\n> ")
			}

			emailAddress = InputText("\n? Alamat email\n> ")
			for emailAddress == "" {
				HandleError(errors.New("Alamat email tidak boleh kosong."))
				emailAddress = InputText("\n? Alamat email\n> ")
			}

			password = InputText("\n? Kata sandi\n> ")
			for password == "" || len(password) < 8 {
				if password == "" {
					HandleError(errors.New("Kata sandi tidak boleh kosong."))
				} else {
					HandleError(errors.New("Kata sandi minimal 8 karakter."))
				}

				password = InputText("\n? Kata sandi\n> ")
			}

			AddServiceAccount(serviceName, userName, emailAddress, password, &serviceAccountsCount)
		case 3:
			fmt.Println("\n=== EDIT AKUN LAYANAN ===")

			if serviceAccountsCount == 0 {
				fmt.Println("\nINFO: Belum ada akun layanan.")
			} else {
				ShowServiceAccounts(serviceName, userName, emailAddress, password, &serviceAccountsCount)

				indexInput = InputText("\n? Nomor akun yang akan diedit\n> ")
				serviceAccountIndex, ok = HandleIndex(indexInput, &serviceAccountsCount)
				for !ok {
					indexInput = InputText("\n? Nomor akun yang akan diedit\n> ")
					serviceAccountIndex, ok = HandleIndex(indexInput, &serviceAccountsCount)
				}

				fmt.Println("\nINFO: Tekan Enter jika tidak ingin mengubah.")

				serviceName = InputText("\n? Nama layanan baru\n> ")
				if serviceName == "" {
					serviceName = ServiceAccounts[serviceAccountIndex-1].NamaLayanan
				}

				userName = InputText("\n? Nama pengguna baru\n> ")
				if userName == "" {
					userName = ServiceAccounts[serviceAccountIndex-1].NamaPengguna
				}

				emailAddress = InputText("\n? Alamat baru\n> ")
				if emailAddress == "" {
					emailAddress = ServiceAccounts[serviceAccountIndex-1].AlamatEmail
				}

				password = InputText("\n? Kata sandi baru\n> ")
				for password != "" && len(password) < 8 {
					HandleError(errors.New("Kata sandi minimal 8 karakter."))
					password = InputText("\n? Kata sandi baru\n> ")
				}
				if password == "" {
					password = ServiceAccounts[serviceAccountIndex-1].KataSandi
				}

				EditServiceAccount(serviceAccountIndex, serviceName, userName, emailAddress, password)
			}
		case 4:
			fmt.Println("\n=== HAPUS AKUN LAYANAN ===")

			if serviceAccountsCount == 0 {
				fmt.Println("\nINFO: Belum ada akun layanan.")
			} else {
				ShowServiceAccounts(serviceName, userName, emailAddress, password, &serviceAccountsCount)

				indexInput = InputText("\n? Nomor akun yang akan dihapus\n> ")
				serviceAccountIndex, ok = HandleIndex(indexInput, &serviceAccountsCount)
				for !ok {
					indexInput = InputText("\n? Nomor akun yang akan dihapus\n> ")
					serviceAccountIndex, ok = HandleIndex(indexInput, &serviceAccountsCount)
				}

				DeleteServiceAccount(serviceAccountIndex, &serviceAccountsCount)
			}
		case 5:
			fmt.Println("\n=== CARI NAMA LAYANAN ===")

			if serviceAccountsCount == 0 {
				fmt.Println("\nINFO: Belum ada akun layanan.")
			} else {
				ShowServiceAccounts(serviceName, userName, emailAddress, password, &serviceAccountsCount)

				fmt.Println("\nMetode pencarian:")
				fmt.Println("\n1. Sequential Search")
				fmt.Println("2. Binary Search")

				seInput = InputText("\n? Pilih metode\n> ")
				se, err = strconv.Atoi(seInput)
				for err != nil {
					HandleError(errors.New("Pilihan harus berupa angka."))

					seInput = InputText("\n? Pilih metode\n> ")
					se, err = strconv.Atoi(seInput)
				}
				fmt.Println()

				switch se {
				case 1:
					fmt.Println("\n=== CARI NAMA LAYANAN (SEQUENTIAL SEARCH) ===")

					keyword = InputText("\n? Kata kunci nama layanan\n> ")
					SeqSearchByServiceName(keyword, &serviceAccountsCount)
				case 2:
					fmt.Println("\n=== CARI NAMA LAYANAN (BINARY SEARCH) ===")

					keyword = InputText("\n? Kata kunci nama layanan\n> ")
				default:
					HandleError(errors.New("Pilihan tidak valid."))
				}
			}
		case 6:
			fmt.Println("\n=== URUTKAN AKUN LAYANAN ===")
		case 0:
			fmt.Println("\n=== SAMPAI JUMPA! ===")
			os.Exit(0)
		default:
			HandleError(errors.New("Pilihan tidak valid."))
		}
	}
}
