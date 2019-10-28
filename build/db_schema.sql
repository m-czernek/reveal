CREATE TABLE IF NOT EXISTS person (
    uuid BINARY(16) PRIMARY KEY,
    years_of_experience DECIMAL(2,1),
    years_at_redhat DECIMAL(2,1),
    updated DATE NOT NULL,
    job_id INT NOT NULL,
    loc_id INT NOT NULL,
    sal_id INT NOT NULL,
    KEY `job_id` (`job_id`),
    CONSTRAINT `fk_job_id` FOREIGN KEY (`job_id`) REFERENCES `job` (`id`),
    KEY `loc_id` (`loc_id`),
    CONSTRAINT `fk_loc_id` FOREIGN KEY (`loc_id`) REFERENCES `location` (`id`),
    KEY `sal_id` (`sal_id`),
    CONSTRAINT `items_idfk_3` FOREIGN KEY (`sal_id`) REFERENCES `salary` (`id`)
);

CREATE TABLE IF NOT EXISTS job (
    id INT AUTO_INCREMENT PRIMARY KEY,
    position_name VARCHAR(50),
    position_type VARCHAR(50),
    position_rank INT CHECK(position_rank > 0 AND position_rank < 10)
);

CREATE TABLE IF NOT EXISTS location (
    id INT AUTO_INCREMENT PRIMARY KEY,
    country VARCHAR(50) NOT NULL,
    city VARCHAR(50),
    remote BOOLEAN
);

CREATE TABLE IF NOT EXISTS salary (
    id INT AUTO_INCREMENT PRIMARY KEY,
    yearly_salary INT NOT NULL,
    currency VARCHAR(10)
);

/* insert into person(uuid, years_of_experience, updated, job_id, loc_id, sal_id) values( UUID_TO_BIN(UUID()), 2.111, DATE(NOW()), 1, 1, 1);
select BIN_TO_UUID(p.uuid) id,p.years_of_experience, p.updated, j.position_name, s.yearly_salary from person p, job j, salary s where p.job_id = j.id and p.sal_id = s.id; */
