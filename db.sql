
CREATE TABLE patient (
    first_name_th VARCHAR(100),
    middle_name_th VARCHAR(100),
    last_name_th VARCHAR(100),
    first_name_en VARCHAR(100),
    middle_name_en VARCHAR(100),
    last_name_en VARCHAR(100),
    date_of_birth DATE,
    patient_hn VARCHAR(50),
    national_id VARCHAR(13) ,
    passport_id VARCHAR(9) ,
    phone_number VARCHAR(20),
    email VARCHAR(100),
    gender VARCHAR(1),    
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT null,
	updated_at TIMESTAMPTZ NULL
);

-- ใช้ UNIQUE INDEX แทน PK
CREATE UNIQUE INDEX idx_patient_national_id ON patient(national_id) WHERE national_id IS NOT NULL;
CREATE UNIQUE INDEX idx_patient_passport_id ON patient(passport_id) WHERE passport_id IS NOT NULL;
CREATE UNIQUE INDEX idx_patient_hn ON patient(patient_hn) WHERE patient_hn IS NOT NULL;
-- สร้าง Indexes สำหรับเพิ่มความเร็วในการค้นหาข้อมูล
CREATE INDEX idx_patient_email ON patient(email);

CREATE TABLE staff (
	id bigserial NOT NULL,
	status int null default(1),
	username varchar(255) NOT NULL,
	hospital varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT null,
	updated_at TIMESTAMPTZ null,
	last_login TIMESTAMPTZ null,
	CONSTRAINT staff_pkey PRIMARY KEY (id)
);

INSERT INTO patient (
    first_name_th,
    middle_name_th,
    last_name_th,
    first_name_en,
    middle_name_en,
    last_name_en,
    date_of_birth,
    patient_hn,
    national_id,
    passport_id,
    phone_number,
    email,
    gender,
    created_at
) VALUES 
-- คนที่ 1: คนไทย (มีเลขบัตรประชาชน ไม่มีพาสปอร์ต)
(
    'สมชาย',
    NULL,
    'ใจดี',
    'Somchai',
    NULL,
    'Jaidee',
    '1985-04-12',
    'HN660000001',
    '1100100200301',
    NULL,
    '0812345678',
    'somchai.j@example.com',
    'M',
    CURRENT_TIMESTAMP
),
-- คนที่ 2: คนไทย (มีชื่อกลาง)
(
    'นภา',
    'สิริ',
    'พานิช',
    'Napa',
    'Siri',
    'Phanich',
    '1992-09-25',
    'HN660000002',
    '1100100200302',
    NULL,
    '0898765432',
    'napa.p@example.com',
    'F',
    CURRENT_TIMESTAMP
),
-- คนที่ 3: ต่างชาติ (มีพาสปอร์ต ไม่มีเลขบัตรประชาชน)
(
    NULL,
    NULL,
    NULL,
    'John',
    'Robert',
    'Smith',
    '1988-11-03',
    'HN660000003',
    NULL,
    'A12345678',
    '0821112233',
    'john.smith@example.com',
    'M',
    CURRENT_TIMESTAMP
),
-- คนที่ 4: ต่างชาติ (มีชื่อกลาง)
(
    NULL,
    NULL,
    NULL,
    'Emma',
    'Louise',
    'Watson',
    '1995-03-30',
    'HN660000004',
    NULL,
    'B98765432',
    '0834445566',
    'emma.w@example.com',
    'F',
    CURRENT_TIMESTAMP
),
-- คนที่ 5: คนไทย (มีทั้งเลขบัตรประชาชน และพาสปอร์ต)
(
    'กิตติ',
    NULL,
    'วัฒนากุล',
    'Kitti',
    NULL,
    'Wattanakul',
    '2000-01-15',
    'HN660000005',
    '1100100200305',
    'C55667788',
    '0847778899',
    'kitti.w@example.com',
    'M',
    CURRENT_TIMESTAMP
);
