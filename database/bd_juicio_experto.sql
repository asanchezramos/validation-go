-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 28-10-2020 a las 12:44:05
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
-- Estructura de tabla para la tabla `criterio`
--

CREATE TABLE `criterio` (
  `criterio_id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `speciality` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `expert_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `criterio`
--

INSERT INTO `criterio` (`criterio_id`, `name`, `speciality`, `expert_id`, `created_at`, `updated_at`) VALUES
(1, 'Criterio 1', 'Software', 12, '2020-10-18 19:30:03', '2020-10-19 00:43:28'),
(2, 'Criterio 1', 'Redes Neuronales', 12, '2020-10-18 19:31:09', '2020-10-19 00:44:23'),
(3, 'Criterio 1', 'Software', 13, '2020-10-18 19:31:32', '2020-10-19 00:44:12'),
(4, 'Criterio 1', 'Redes Neuronales', 13, '2020-10-18 19:31:32', '2020-10-19 00:44:27'),
(5, 'criterio 2', 'Software', 12, '2020-10-20 16:28:50', '2020-10-20 16:28:50'),
(6, 'criterio 3', 'Software', 12, '2020-10-20 16:28:50', '2020-10-20 16:28:50');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `criterio_response`
--

CREATE TABLE `criterio_response` (
  `criterio_response_id` int(11) NOT NULL,
  `criterio_id` int(11) DEFAULT NULL,
  `research_id` int(11) DEFAULT NULL,
  `dimension_id` int(11) NOT NULL,
  `status` enum('A','R','D') COLLATE utf8_spanish_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `criterio_response`
--

INSERT INTO `criterio_response` (`criterio_response_id`, `criterio_id`, `research_id`, `dimension_id`, `status`, `created_at`, `updated_at`) VALUES
(25, 1, 31, 56, 'A', '2020-10-21 21:57:59', '2020-10-21 21:57:59'),
(26, 5, 31, 56, 'R', '2020-10-21 21:58:02', '2020-10-21 21:58:02'),
(27, 6, 31, 56, 'A', '2020-10-21 21:58:05', '2020-10-21 21:58:05'),
(28, 1, 31, 52, 'R', '2020-10-21 21:58:20', '2020-10-21 21:58:20'),
(29, 5, 31, 52, 'A', '2020-10-21 21:58:21', '2020-10-21 21:58:21'),
(30, 6, 31, 52, 'A', '2020-10-21 21:58:22', '2020-10-21 21:58:56'),
(31, 5, 31, 54, 'A', '2020-10-21 21:58:25', '2020-10-21 21:58:25'),
(32, 6, 31, 54, 'R', '2020-10-21 21:58:32', '2020-10-21 21:58:32'),
(33, 1, 31, 54, 'A', '2020-10-21 21:58:36', '2020-10-21 21:58:36'),
(34, 1, 31, 55, 'A', '2020-10-21 21:58:42', '2020-10-21 21:58:42'),
(35, 5, 31, 55, 'R', '2020-10-21 21:58:47', '2020-10-21 21:58:47'),
(36, 6, 31, 53, 'A', '2020-10-21 21:59:02', '2020-10-21 21:59:02'),
(37, 5, 31, 53, 'A', '2020-10-21 21:59:04', '2020-10-21 21:59:04'),
(38, 6, 31, 55, 'A', '2020-10-21 21:59:07', '2020-10-21 21:59:07'),
(39, 1, 31, 53, 'R', '2020-10-21 21:59:11', '2020-10-21 21:59:11'),
(40, 1, 35, 63, 'A', '2020-10-22 00:30:19', '2020-10-22 00:44:04'),
(41, 5, 35, 63, 'A', '2020-10-22 00:30:22', '2020-10-22 00:33:18'),
(42, 6, 35, 63, 'A', '2020-10-22 00:30:23', '2020-10-22 00:30:23'),
(43, 1, 35, 62, 'A', '2020-10-22 00:32:35', '2020-10-22 00:32:35'),
(44, 5, 35, 62, 'A', '2020-10-22 00:32:39', '2020-10-22 00:32:39'),
(45, 6, 35, 62, 'A', '2020-10-22 00:32:42', '2020-10-22 00:32:42'),
(46, 1, 35, 64, 'A', '2020-10-22 00:32:47', '2020-10-22 00:32:47'),
(47, 5, 35, 64, 'A', '2020-10-22 00:32:53', '2020-10-22 00:33:21'),
(48, 6, 35, 64, 'A', '2020-10-22 00:32:56', '2020-10-22 00:33:23'),
(49, 1, 36, 66, 'A', '2020-10-24 21:13:04', '2020-10-24 21:13:04'),
(50, 5, 36, 66, 'A', '2020-10-24 21:13:06', '2020-10-24 21:13:06'),
(51, 6, 36, 66, 'R', '2020-10-24 21:13:08', '2020-10-24 21:16:38'),
(52, 1, 40, 69, 'A', '2020-10-25 16:37:36', '2020-10-25 16:37:36'),
(53, 5, 40, 69, 'A', '2020-10-25 16:37:38', '2020-10-28 04:45:58'),
(54, 6, 40, 69, 'R', '2020-10-25 16:37:39', '2020-10-25 16:37:39'),
(55, 1, 40, 70, 'A', '2020-10-25 16:48:30', '2020-10-25 16:48:30'),
(56, 5, 40, 70, 'A', '2020-10-25 16:48:31', '2020-10-25 16:48:31'),
(57, 6, 40, 70, 'A', '2020-10-25 16:48:32', '2020-10-25 16:48:32'),
(58, 1, 41, 71, 'A', '2020-10-28 04:43:39', '2020-10-28 04:43:39'),
(59, 5, 41, 71, 'A', '2020-10-28 04:43:40', '2020-10-28 04:43:40'),
(60, 6, 41, 71, 'A', '2020-10-28 04:43:40', '2020-10-28 04:43:40'),
(61, 1, 40, 72, 'A', '2020-10-28 04:50:12', '2020-10-28 04:50:12'),
(62, 5, 40, 72, 'A', '2020-10-28 04:50:12', '2020-10-28 04:50:12'),
(63, 6, 40, 72, 'A', '2020-10-28 04:50:13', '2020-10-28 04:50:13'),
(64, 1, 44, 76, 'A', '2020-10-28 04:59:42', '2020-10-28 04:59:42'),
(65, 5, 44, 76, 'R', '2020-10-28 04:59:43', '2020-10-28 04:59:43'),
(66, 6, 44, 76, 'R', '2020-10-28 04:59:43', '2020-10-28 04:59:43'),
(67, 1, 48, 81, 'A', '2020-10-28 06:45:42', '2020-10-28 06:45:42'),
(68, 5, 48, 81, 'A', '2020-10-28 06:45:43', '2020-10-28 06:45:43'),
(69, 6, 48, 81, 'A', '2020-10-28 06:45:44', '2020-10-28 06:45:44'),
(70, 1, 48, 84, 'A', '2020-10-28 06:45:45', '2020-10-28 06:45:45'),
(71, 5, 48, 84, 'A', '2020-10-28 06:45:47', '2020-10-28 06:45:47'),
(72, 6, 48, 84, 'A', '2020-10-28 06:45:48', '2020-10-28 06:45:48'),
(73, 1, 48, 82, 'A', '2020-10-28 06:45:51', '2020-10-28 06:45:51'),
(74, 6, 48, 83, 'A', '2020-10-28 06:45:54', '2020-10-28 06:45:54'),
(75, 1, 48, 83, 'A', '2020-10-28 06:45:56', '2020-10-28 06:45:56'),
(76, 5, 48, 83, 'A', '2020-10-28 06:45:58', '2020-10-28 06:45:58'),
(77, 5, 48, 82, 'A', '2020-10-28 06:46:00', '2020-10-28 06:46:00'),
(78, 6, 48, 82, 'A', '2020-10-28 06:46:02', '2020-10-28 06:46:02');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `dimension`
--

CREATE TABLE `dimension` (
  `dimension_id` int(11) NOT NULL,
  `research_id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `variable` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `status` enum('P','A','R') COLLATE utf8_spanish_ci NOT NULL DEFAULT 'P',
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `dimension`
--

INSERT INTO `dimension` (`dimension_id`, `research_id`, `name`, `variable`, `status`, `created_at`, `updated_at`) VALUES
(52, 31, 'Tiempo	', 'Tiempo de respuesta de las autoridades policiales', 'P', '2020-10-21 19:20:23', '2020-10-21 19:20:23'),
(53, 31, 'Cantidad', 'Cantidad de denuncias por maltrato psicologico', 'P', '2020-10-21 19:20:47', '2020-10-21 19:20:47'),
(54, 31, 'Tiempo	', 'Tiempo encomendado para la lenvantacion de informacion', 'P', '2020-10-21 19:21:06', '2020-10-21 19:21:06'),
(55, 31, 'Cantidad', 'Cantidad de datos recopilados por usuario para poder activar la ayuda', 'P', '2020-10-21 19:21:23', '2020-10-21 19:21:23'),
(56, 31, 'Tiempo	', 'Tiempo de medida en las altas esferas', 'P', '2020-10-21 19:21:35', '2020-10-21 19:21:35'),
(57, 31, 'Cantidad', 'asd', 'P', '2020-10-21 22:49:51', '2020-10-21 22:49:51'),
(58, 31, 'Cantidad', 'asd', 'P', '2020-10-21 22:50:36', '2020-10-21 22:50:36'),
(59, 32, 'asdasd', 'asdasdasd', 'P', '2020-10-21 23:31:13', '2020-10-21 23:31:13'),
(60, 32, 'aaaaaaaaa', 'asd', 'P', '2020-10-21 23:31:53', '2020-10-21 23:31:53'),
(62, 35, 'tiempo	', 'asdasd asd asd asd ', 'P', '2020-10-22 00:26:00', '2020-10-22 00:26:00'),
(63, 35, 'tiempo	', 'asdas das dasd as das d', 'P', '2020-10-22 00:26:38', '2020-10-22 00:26:38'),
(64, 35, 'asdasdasdasd', 'asd asda sdas d', 'P', '2020-10-22 00:27:03', '2020-10-22 00:27:03'),
(65, 38, 'jjj', 'nnmklk', 'P', '2020-10-24 20:36:23', '2020-10-24 20:36:23'),
(66, 36, 'dffvg', 'gghhh', 'P', '2020-10-24 21:07:06', '2020-10-24 21:07:06'),
(67, 39, 'asd', 'asd 213 123 123 ', 'P', '2020-10-25 03:55:41', '2020-10-25 05:17:31'),
(68, 39, 'asd', 'sd', 'P', '2020-10-25 03:56:09', '2020-10-25 03:56:09'),
(71, 41, 'tiempo	', 'cantidad de tiempo', 'P', '2020-10-28 04:41:49', '2020-10-28 04:41:49'),
(72, 40, 'sdfsdf', 'sdfsdfsdf', 'P', '2020-10-28 04:47:14', '2020-10-28 04:47:14'),
(73, 42, 'asd', 'asd', 'P', '2020-10-28 04:49:46', '2020-10-28 04:49:46'),
(74, 43, 'asd', 'asd', 'P', '2020-10-28 04:54:50', '2020-10-28 04:54:50'),
(75, 43, 'asd', 'asd', 'P', '2020-10-28 04:57:25', '2020-10-28 04:57:25'),
(76, 44, 's', 's asdasd asd asd asd ', 'P', '2020-10-28 04:59:17', '2020-10-28 05:37:54'),
(77, 45, 'asd', 'asd', 'P', '2020-10-28 05:38:22', '2020-10-28 05:38:22'),
(78, 46, 'ss', 'ssssss', 'P', '2020-10-28 05:41:18', '2020-10-28 05:41:24'),
(79, 46, 'asd', 'asd', 'P', '2020-10-28 05:41:32', '2020-10-28 05:41:32'),
(80, 47, 'sdf', 'sdf', 'P', '2020-10-28 06:35:46', '2020-10-28 06:35:46'),
(81, 48, 'sdf', 'sdf', 'P', '2020-10-28 06:39:23', '2020-10-28 06:39:23'),
(82, 48, 'asd', 'asd', 'P', '2020-10-28 06:39:36', '2020-10-28 06:39:36'),
(83, 48, 'asd', 'asd', 'P', '2020-10-28 06:42:00', '2020-10-28 06:42:00'),
(84, 48, 'sdf', 'sss', 'P', '2020-10-28 06:42:10', '2020-10-28 06:42:10');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `network`
--

CREATE TABLE `network` (
  `network_id` int(11) NOT NULL,
  `user_base_id` int(11) NOT NULL,
  `user_relation_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `network`
--

INSERT INTO `network` (`network_id`, `user_base_id`, `user_relation_id`, `created_at`, `updated_at`) VALUES
(6, 11, 12, '2020-10-28 04:53:59', '2020-10-28 04:53:59'),
(7, 11, 13, '2020-10-28 06:33:25', '2020-10-28 06:33:25'),
(8, 11, 14, '2020-10-28 06:33:25', '2020-10-28 06:33:25');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `network_request`
--

CREATE TABLE `network_request` (
  `network_request_id` int(11) NOT NULL,
  `user_base_id` int(11) NOT NULL,
  `user_relation_id` int(11) NOT NULL,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `network_request`
--

INSERT INTO `network_request` (`network_request_id`, `user_base_id`, `user_relation_id`, `status`, `created_at`, `updated_at`) VALUES
(20, 11, 12, 2, '2020-10-28 04:53:40', '2020-10-28 04:53:59');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `notification`
--

CREATE TABLE `notification` (
  `notification_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `title` varchar(100) COLLATE utf8_spanish_ci DEFAULT NULL,
  `body` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `research`
--

CREATE TABLE `research` (
  `research_id` int(11) NOT NULL,
  `researcher_id` int(11) NOT NULL,
  `expert_id` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `speciality` varchar(255) COLLATE utf8_spanish_ci NOT NULL,
  `authors` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `observation` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `attachment_one` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `attachment_two` varchar(255) COLLATE utf8_spanish_ci DEFAULT NULL,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `research`
--

INSERT INTO `research` (`research_id`, `researcher_id`, `expert_id`, `title`, `speciality`, `authors`, `observation`, `attachment_one`, `attachment_two`, `status`, `created_at`, `updated_at`) VALUES
(48, 11, 12, 'sdf', 'Software', 'sdf', 'ok', '1603867158-IMG_20201019_061510.jpg', '1603867158-IMG_20201019_061510.jpg', 3, '2020-10-28 06:39:18', '2020-10-28 06:46:08');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `resource_user`
--

CREATE TABLE `resource_user` (
  `resource_user_id` int(11) NOT NULL,
  `expert_id` int(11) DEFAULT NULL,
  `title` varchar(300) COLLATE utf8_spanish_ci DEFAULT NULL,
  `subtitle` varchar(400) COLLATE utf8_spanish_ci DEFAULT NULL,
  `link` varchar(500) COLLATE utf8_spanish_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

--
-- Volcado de datos para la tabla `resource_user`
--

INSERT INTO `resource_user` (`resource_user_id`, `expert_id`, `title`, `subtitle`, `link`, `created_at`, `updated_at`) VALUES
(2, 0, 'sdfghj', 'dfghjklñ', 'https://flutter.dev/?gclid=Cj0KCQjwit_8BRCoARIsAIx3Rj7dh-KYRFxCP11JFceAWDcntZBQbp-LiWpV-bv0jkRN49BCTWsMX5IaAg0SEALw_wcB&gclsrc=aw.ds', '2020-10-27 13:09:36', '2020-10-28 06:11:33'),
(3, 0, 'sdfghj', 'dfghjklñ', 'https://flutter.dev/?gclid=Cj0KCQjwit_8BRCoARIsAIx3Rj7dh-KYRFxCP11JFceAWDcntZBQbp-LiWpV-bv0jkRN49BCTWsMX5IaAg0SEALw_wcB&gclsrc=aw.ds', '2020-10-27 13:10:35', '2020-10-28 06:11:37'),
(7, 0, 'tiruko 12', 'asdasd', 'https://github.com/tjcrisostomob/ApiRestLibMerceMily/settings/access', '2020-10-28 06:13:46', '2020-10-28 06:13:46'),
(8, 0, 'asd', 'asd', 'asd', '2020-10-28 06:15:17', '2020-10-28 06:15:17'),
(9, 0, 'asd', 'asd', 'asdasdasd', '2020-10-28 06:15:29', '2020-10-28 06:15:29'),
(11, 12, 'asd', 'asd', 'asdasdasd', '2020-10-28 06:17:37', '2020-10-28 06:17:37'),
(12, 12, 'dsf', 'sdf', 'sdf', '2020-10-28 06:24:28', '2020-10-28 06:24:28'),
(13, 12, 'asd', 'asd', 'asd', '2020-10-28 06:27:33', '2020-10-28 06:27:33'),
(14, 12, 'asd', 'asd', 'asd', '2020-10-28 06:27:37', '2020-10-28 06:27:37'),
(15, 12, 'asda', 'sd', 'asd', '2020-10-28 06:27:40', '2020-10-28 06:27:40'),
(16, 12, 'sss', 'ss', 'ss', '2020-10-28 06:27:45', '2020-10-28 06:27:45');

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
(11, 'Investigador', 'Usuario', '1603661312-image_picker1494451654793012890.jpg', 'asd@asd', '$2a$10$l5xpy76AD7rWTNFnAZwpyO72NLkVtPL.4LQsyPhqZHeMT.agGsfH6', '', NULL, 'U', 1, '2020-10-16 01:24:51', '2020-10-28 06:32:40'),
(12, 'Experto Primaria', 'Primero aaaa', '1603661312-image_picker1494451654793012890.jpg', 'qwe@qwe', '$2a$10$0c.UkJQjfK2b4YPLNKuEX.50d70J0mX/wwh7D3bsAzk.4SNi8fEJG', '12336524', 'Software', 'E', 1, '2020-10-16 01:43:50', '2020-10-28 06:32:38'),
(13, 'Experto', 'Segundo', '1603661312-image_picker1494451654793012890.jpg', 'ert@ert', '$2a$10$enJy3s8zXqtd2PbxddoarOox.Cnz7vsc4RUMjTnEZk2l1sG7zD94O', '956894668', 'Redes Neuronales', 'E', 1, '2020-10-16 03:35:31', '2020-10-28 06:32:36'),
(14, 'junior	teofilo', 'berrocal condori', '1603661312-image_picker1494451654793012890.jpg', 'aaa@aaa', '$2a$10$aKbTdEgw.FUT3o.ZfKcfuOU7AOZCiQrF34Ff8PbB0vzYhEs4gUx1q', '956894668', 'Software', 'E', 1, '2020-10-25 21:28:32', '2020-10-25 21:28:32'),
(15, 'Teofilo', 'Crisostomo', '1603661312-image_picker1494451654793012890.jpg', 'tjcrisostomob@gmail.com', '$2a$10$GabQeaauKBfpLB7v0uRHp.B55ZpplB7pcrE/rfcKylUCP08WidNQu', '956894668', 'Software', 'U', 1, '2020-10-27 13:04:41', '2020-10-28 06:32:43');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `criterio`
--
ALTER TABLE `criterio`
  ADD PRIMARY KEY (`criterio_id`);

--
-- Indices de la tabla `criterio_response`
--
ALTER TABLE `criterio_response`
  ADD PRIMARY KEY (`criterio_response_id`),
  ADD UNIQUE KEY `criterio_response_id` (`criterio_response_id`,`criterio_id`,`research_id`,`dimension_id`);

--
-- Indices de la tabla `dimension`
--
ALTER TABLE `dimension`
  ADD PRIMARY KEY (`dimension_id`);

--
-- Indices de la tabla `network`
--
ALTER TABLE `network`
  ADD PRIMARY KEY (`network_id`);

--
-- Indices de la tabla `network_request`
--
ALTER TABLE `network_request`
  ADD PRIMARY KEY (`network_request_id`),
  ADD UNIQUE KEY `user_base_id` (`user_base_id`,`user_relation_id`);

--
-- Indices de la tabla `research`
--
ALTER TABLE `research`
  ADD PRIMARY KEY (`research_id`);

--
-- Indices de la tabla `resource_user`
--
ALTER TABLE `resource_user`
  ADD PRIMARY KEY (`resource_user_id`);

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
-- AUTO_INCREMENT de la tabla `criterio`
--
ALTER TABLE `criterio`
  MODIFY `criterio_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT de la tabla `criterio_response`
--
ALTER TABLE `criterio_response`
  MODIFY `criterio_response_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=79;

--
-- AUTO_INCREMENT de la tabla `dimension`
--
ALTER TABLE `dimension`
  MODIFY `dimension_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=85;

--
-- AUTO_INCREMENT de la tabla `network`
--
ALTER TABLE `network`
  MODIFY `network_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT de la tabla `network_request`
--
ALTER TABLE `network_request`
  MODIFY `network_request_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT de la tabla `research`
--
ALTER TABLE `research`
  MODIFY `research_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=49;

--
-- AUTO_INCREMENT de la tabla `resource_user`
--
ALTER TABLE `resource_user`
  MODIFY `resource_user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT de la tabla `user`
--
ALTER TABLE `user`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
