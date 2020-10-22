-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 15-10-2020 a las 15:49:36
-- Versión del servidor: 10.4.14-MariaDB
-- Versión de PHP: 7.2.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `bd_juicio_experto`
--
CREATE DATABASE IF NOT EXISTS `bd_juicio_experto` DEFAULT CHARACTER SET utf8 COLLATE utf8_spanish_ci;
USE `bd_juicio_experto`;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `answer`
--

CREATE TABLE `answer` (
  `answer_id` int(11) NOT NULL,
  `comments` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `file` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `solicitude_id` int(11) NOT NULL,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `solicitude`
--

CREATE TABLE `solicitude` (
  `solicitude_id` int(11) NOT NULL,
  `repository` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `investigation` varchar(100) COLLATE utf8_spanish_ci DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `expert_id` int(11) NOT NULL,
  `status` char(1) COLLATE utf8_spanish_ci DEFAULT 'P',
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `user`
--

CREATE TABLE `user` (
  `user_id` int(11) NOT NULL,
  `name` varchar(100) COLLATE utf8_spanish_ci NOT NULL,
  `full_name` varchar(100) COLLATE utf8_spanish_ci NOT NULL,
  `photo` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `mail` varchar(45) COLLATE utf8_spanish_ci NOT NULL,
  `password` varchar(100) COLLATE utf8_spanish_ci NOT NULL,
  `phone` varchar(15) COLLATE utf8_spanish_ci NOT NULL,
  `specialty` varchar(100) COLLATE utf8_spanish_ci DEFAULT NULL,
  `role` enum('U','E') COLLATE utf8_spanish_ci NOT NULL DEFAULT 'U',
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `user`
--

INSERT INTO `user` (`user_id`, `name`, `full_name`, `photo`, `mail`, `password`, `phone`, `specialty`, `role`, `status`, `created_at`, `updated_at`) VALUES
(3, 'Teofilo', 'Crisostomo', NULL, 'tjcrisostomob@gmail.com', '$2a$10$oRL1.ucTI3dnrLzFuytvPe7nXpWFUkn5Vy2eCzqN2kxg6hLEq6v.i', '956894668', 'Software', 'U', 1, '2020-10-14 13:23:43', '2020-10-14 13:23:43'),
(4, 'Crisdótian', 'Torres', NULL, 'ju.cri.xd@gmail.com', '$2a$10$0bixa2SwcyL8v3Z.wX/ptuJSMgWAdQpuwY.fifvy6i/Jkixcx2s8O', '956894668', 'Software', 'E', 1, '2020-10-14 13:37:32', '2020-10-15 00:55:05'),
(5, 'Cristian', 'Torres', NULL, 'ju.cri.xd@gmail.com', '$2a$10$6ROf1H3Nxk6MBvgNnOlGHevgnk/2bQskeTu00CI/inCyNz64ClHsy', '956894668', 'Software', 'E', 1, '2020-10-15 03:50:27', '2020-10-15 03:50:27'),
(6, 'sdasd', 'asdasd', NULL, 'asdads@asdasd.com', '$2a$10$OacggLWQowzKh0SE1EOBneQV6hRUY8M7sO6m/rf0zy32ZEDl0moTe', '55666666', 'Redes Neuronales', 'E', 1, '2020-10-15 03:53:11', '2020-10-15 03:53:11'),
(7, 'Teófilo Junior', 'Crisóstomo Berrocal', '1602737677-image_picker4811691699311310173.jpg', 'tjcrisostomob@outlook.com', '$2a$10$et/.VKthe2lBVLDHK1t3q.i8b6ltdAcO9UzUh530vwDwQbMWFRz9e', '956894668', 'Software', 'E', 1, '2020-10-15 04:54:37', '2020-10-15 04:54:37');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `answer`
--
ALTER TABLE `answer`
  ADD PRIMARY KEY (`answer_id`),
  ADD KEY `FK_answer_solicitude_idx` (`solicitude_id`),
  ADD KEY `IDX_answer_status` (`status`);

--
-- Indices de la tabla `solicitude`
--
ALTER TABLE `solicitude`
  ADD PRIMARY KEY (`solicitude_id`),
  ADD KEY `FK_solicitude_user_idx` (`user_id`),
  ADD KEY `FK_solicitude_expert_idx` (`expert_id`),
  ADD KEY `IDX_solicitude_status` (`status`);

--
-- Indices de la tabla `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`user_id`),
  ADD KEY `IDX_user_status` (`status`),
  ADD KEY `IDX_user_role` (`role`),
  ADD KEY `IDX_user_email` (`mail`),
  ADD KEY `IDX_user_password` (`password`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `answer`
--
ALTER TABLE `answer`
  MODIFY `answer_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `solicitude`
--
ALTER TABLE `solicitude`
  MODIFY `solicitude_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `user`
--
ALTER TABLE `user`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `answer`
--
ALTER TABLE `answer`
  ADD CONSTRAINT `FK_answer_solicitude` FOREIGN KEY (`solicitude_id`) REFERENCES `solicitude` (`solicitude_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Filtros para la tabla `solicitude`
--
ALTER TABLE `solicitude`
  ADD CONSTRAINT `FK_solicitude_expert` FOREIGN KEY (`expert_id`) REFERENCES `user` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `FK_solicitude_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
