-- 000003_cms_tables.up.sql

-- Drop old tables that we want to restructure
DROP TABLE IF EXISTS education CASCADE;
DROP TABLE IF EXISTS experiences CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;

-- Create New Hero Section Table
CREATE TABLE IF NOT EXISTS hero_section (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    headline TEXT NOT NULL,
    subheadline TEXT,
    biography TEXT,
    profile_image TEXT,
    resume_url TEXT,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create New Skills Table
CREATE TABLE IF NOT EXISTS skills (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    icon_class TEXT,
    category TEXT NOT NULL,
    display_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create New Social Links Table
CREATE TABLE IF NOT EXISTS social_links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    platform TEXT NOT NULL,
    url TEXT NOT NULL,
    icon_name TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create New Experience Table (Singular and better dates)
CREATE TABLE IF NOT EXISTS experience (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company TEXT NOT NULL,
    role TEXT NOT NULL,
    description TEXT,
    start_date DATE NOT NULL,
    end_date DATE,
    is_current BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create New Education Table (Better dates)
CREATE TABLE IF NOT EXISTS education (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    institution TEXT NOT NULL,
    degree TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Seed initial hero data
INSERT INTO hero_section (headline, subheadline, biography, profile_image, resume_url)
VALUES (
    'Building reliable web apps with a focus on quality.',
    'I''m a Full-Stack Developer & QA Specialist',
    'I''m a Full-Stack Developer & QA Specialist with over 3 years of experience in Angular and Firebase, building functional web applications and ensuring quality through rigorous testing.',
    'assets/profile2.jpg',
    'assets/facundo-salvucci_cv.pdf'
);

-- Seed initial social links (larger icons)
INSERT INTO social_links (platform, url, icon_name) VALUES 
('GitHub', 'https://github.com/SalvucciFacundo', 'github'),
('LinkedIn', 'https://www.linkedin.com/in/facundo-salvucci', 'linkedin');
