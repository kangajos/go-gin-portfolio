package models

type Portfolio struct {
	Profile   Profile
	Skill     []Skill
	Social    Social
	Education []Education
	Project   []Project
}

func (p *Portfolio) Get() *Portfolio {
	profile := Profile{}
	p.Profile = *profile.Get()

	skill := Skill{}
	p.Skill = skill.Get()

	social := Social{}
	p.Social = social.Get()

	education := Education{}
	p.Education = education.Get()

	project := Project{}
	p.Project = project.Get()
	return p
}

type Profile struct {
	Name     string
	Email    string
	Phone    string
	Degree   string
	Address  string
	Age      int16
	Position string
	About    string
}

func (p *Profile) Get() *Profile {
	p.Name = "Agus Salim"
	p.Email = "aguss6460@gmail.com"
	p.Phone = "+6281324275549"
	p.Degree = "Bachelor"
	p.Address = "Lumajang, Jawa Timur"
	p.Age = 27
	p.Position = "Backend Developer"
	p.About = "I am Agus Salim, since 2014 I have been studying Informatics Engineering at Muhammiyah University Jember and graduated in 2018. During my college days I started to get to know the world of Web Programming. With the skills I have, I ventured as a Freelancer to hone my skills and in 2018 I was accepted at a company and started my career as a Back End Developer (BE)."
	return p
}

type Education struct {
	Name              string
	LengthOfEducation string
	Address           string
	Description       []string
}

func (e *Education) Get() []Education {
	educations := []Education{
		{Name: "Universitas Muhammadiyah Jember - Teknik Informatika", LengthOfEducation: "2014 - 2018", Address: "Jember, Jawa Timur", Description: []string{"Mengikuti Kompetisi Database - 2016", "Mengikutiti Pelatihan Miktrotik MTCNA - 2017", "Aktif dalam organisasi HUMANIKA & IMM"}},
		{Name: "SMK Miftahul Islam Kunir - Teknik Komputer & Jaringan ", LengthOfEducation: "2011 - 2014", Address: "Kabupaten Lumajang, Jawa Timur", Description: []string{"Mengikuti Anggota OSIS - 2011 s.d 2013"}},
	}
	return educations
}

type Skill struct {
	Stack string
	Range int16
}

func (s *Skill) Get() []Skill {
	skills := []Skill{
		{Stack: "Go Language", Range: 100},
		{Stack: "Go Gin Gonic", Range: 100},
		{Stack: "Go Fiber", Range: 100},
		{Stack: "NodeJS", Range: 100},
		{Stack: "NestJS", Range: 100},
		{Stack: "ExpressJS", Range: 100},
		{Stack: "PHP", Range: 100},
		{Stack: "Laravel", Range: 100},
		{Stack: "Lumen", Range: 100},
		{Stack: "Code Igniter 3", Range: 100},
		{Stack: "MySQL", Range: 100},
		{Stack: "PostgreSQL", Range: 100},
		{Stack: "Docker (basic)", Range: 100},
		{Stack: "GIT Repository Gitlab/Github", Range: 100},
		{Stack: "Java Script", Range: 100},
		{Stack: "JQuery", Range: 100},
		{Stack: "Bootstrap CSS 4/5", Range: 100},
	}
	return skills
}

type Social struct {
	Website  string
	Linkedin string
	Github   string
	Gitlab   string
}

func (s *Social) Get() Social {
	social := Social{
		Website:  "https://www.ajos.com",
		Linkedin: "https://www.linkedin.com/in/agus-salim-590314b1690/?originalSubdomain=id",
		Github:   "https://github.com/kangajos",
		Gitlab:   "https://gitlab.com/agus-salim",
	}
	return social
}

type Project struct {
	Company         string
	LengthOfWorking string
	Position        string
	Description     string
	Detail          []ProjectDetail
}

func (p *Project) Get() []Project {
	projects := []Project{
		{
			Company: "PT. Ako Media Asia", LengthOfWorking: "2022 - Sekarang", Position: "Backend Developer",
			Description: "Sebagai Backend Developer",
			Detail: []ProjectDetail{
				{Name: "Membuat HTML Builder - PT. Pegadaian", Stack: "Laravel 8 Livewire, Bootstrap 5, Java Script, PostgreSQL, Nginx"},
				{Name: "Membuat Chatbot Transaksi - Unilever Indonesia", Stack: "Laravel 8, Redis, Nginx, intragasi with : Wit.ai, Qiscus, Midtrans & Paxel Box"},
				{Name: "Membuat Gadai Smartlocker -  PT. Pegadaian", Stack: "Lumen 8 as Microservices, MySQL, Docker, Integrate wit : Core Sistem, Smartlocker Paxel Box"},
			},
		},
		{
			Company: "PT. Bonafitek Anjar Patria", LengthOfWorking: "2021 - 2022", Position: "Backend Developer",
			Description: "Sebagai Backend Developer",
			Detail: []ProjectDetail{
				{Name: "Mengembangkan Sistem Keuangan Kemenpar dengan Syncrho", Stack: "Symfony 4, Template Admin, JQuery, PostgreSQL, Nginx"},
				{Name: "Membuat Sistem Dokumen Imigrasi", Stack: "Symfony 5, Template Landing Page, JQuery, MySQL & Nginx"},
				{Name: "Mengembangkan Siste Cekal Imigrasi", Stack: "Symfony 5, Template Admin,JQuery ,PostgreSQL & Nginx"},
				{Name: "Revam Web Company Profile ", Stack: "Code Igniter 3, MySQL & Nginx"},
			},
		},
		{
			Company: "PT. Arminareka Pharmasia Pratama", LengthOfWorking: "2020 - 2021", Position: "Backend Developer",
			Description: "Membuat dan mengembangkan sistem internal perusahaan yg bergerak di bidang travel Umroh dan Haji",
			Detail: []ProjectDetail{
				{Name: "Mengembangkan Sistem Mangemen Travel Umroh & Haji", Stack: "Laravel 7, Template AdminLTE3, JQuery, PostgreSQL, Nginx & AWS Light Shail (VPS)"},
				{Name: "Membuat Web Company Profile", Stack: "Code Igniter 3, Template Landing Page, JQuery, MySQL, Nginx & AWS Light Shail (VPS)"},
				{Name: "Mengembangkan Sistem Multi Level Marketing (MLM)", Stack: "Laravel 8, Template Admin,JQuery ,PostgreSQL, Nginx & AWS Light Shail (VPS)"},
				{Name: "Membuat Service Server File (API)", Stack: "Laravel 8, PostgreSQL, Nginx & AWS Light Shail (VPS)"},
			},
		},
		{
			Company: "PT. Jagongan Bisnis", LengthOfWorking: "2018 - 2020", Position: "Backend Developer",
			Description: "Backend Developer",
			Detail: []ProjectDetail{
				{Name: "Sistem Assesmen Pegawai", Stack: "Code Igniter 3, Template AdminLTE, JQuery, MySQL & Apache2"},
				{Name: "Sistem Sertifikasi Karyawan", Stack: "Code Igniter 3, Template Eazy Mobike, MySQL, JQuery & Apache2"},
				{Name: "Sistem Tanda Tangan Elektronik", Stack: "Code Igniter 3, Template AdminLTE, JQuery, MySQL, Apache2 dan 3party BSRE"},
			},
		},
	}
	return projects
}

type ProjectDetail struct {
	Name  string
	Stack string
}
