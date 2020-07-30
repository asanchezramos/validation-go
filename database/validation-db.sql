-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema validation-db
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `validation-db` ;

-- -----------------------------------------------------
-- Schema validation-db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `validation-db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `validation-db` ;

-- -----------------------------------------------------
-- Table `validation-db`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `validation-db`.`user` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `full_name` VARCHAR(100) NOT NULL,
  `photo` VARCHAR(255) NULL,
  `mail` VARCHAR(45) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  `phone` VARCHAR(15) NOT NULL,
  `specialty` VARCHAR(100) NULL,
  `role` ENUM('U', 'E') NOT NULL DEFAULT 'U',
  `status` INT(1) NULL DEFAULT '1',
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  INDEX `IDX_user_status` (`status` ASC) VISIBLE,
  INDEX `IDX_user_role` (`role` ASC) VISIBLE,
  INDEX `IDX_user_email` (`mail` ASC) VISIBLE,
  INDEX `IDX_user_password` (`password` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `validation-db`.`solicitude`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `validation-db`.`solicitude` (
  `solicitude_id` INT NOT NULL AUTO_INCREMENT,
  `repository` VARCHAR(255) NOT NULL,
  `investigation` VARCHAR(100) NULL,
  `user_id` INT NOT NULL,
  `expert_id` INT NOT NULL,
  `status` CHAR(1) NULL DEFAULT 'P',
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`solicitude_id`),
  INDEX `FK_solicitude_user_idx` (`user_id` ASC) VISIBLE,
  INDEX `FK_solicitude_expert_idx` (`expert_id` ASC) VISIBLE,
  INDEX `IDX_solicitude_status` (`status` ASC) VISIBLE,
  CONSTRAINT `FK_solicitude_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `validation-db`.`user` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `FK_solicitude_expert`
    FOREIGN KEY (`expert_id`)
    REFERENCES `validation-db`.`user` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `validation-db`.`answer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `validation-db`.`answer` (
  `answer_id` INT NOT NULL AUTO_INCREMENT,
  `comments` VARCHAR(255) NOT NULL,
  `file` VARCHAR(255) NULL,
  `solicitude_id` INT NOT NULL,
  `status` INT(1) NULL DEFAULT '1',
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`answer_id`),
  INDEX `FK_answer_solicitude_idx` (`solicitude_id` ASC) VISIBLE,
  INDEX `IDX_answer_status` (`status` ASC) VISIBLE,
  CONSTRAINT `FK_answer_solicitude`
    FOREIGN KEY (`solicitude_id`)
    REFERENCES `validation-db`.`solicitude` (`solicitude_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
