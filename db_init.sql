-- phpMyAdmin SQL Dump
-- version 4.9.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 22, 2021 at 05:11 AM
-- Server version: 5.7.34
-- PHP Version: 7.3.27-9+0~20210227.82+debian9~1.gbpa4a3d6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `grocery`
--

-- --------------------------------------------------------

--
-- Table structure for table `shoppinglist`
--

CREATE TABLE `shoppinglist` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(500) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `unit` varchar(500) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `shoppinglist`
--

INSERT INTO `shoppinglist` (`id`, `name`, `qty`, `unit`) VALUES
(2, 'ertyery', 233, 'wertwere'),
(3, 'ERtyer', 12, 'erteey'),
(4, 'BNDSRQ', 123, 'FGDFH'),
(5, 'H%ERHWER', 134, 'HJDATW'),
(9, 'ERtyer123412341234', 1234, 'GFKHHyretreretert'),
(11, 'ASDFASFDaFDGDHDHHHHH', 12345, 'mTYKRTY-+_$WT'),
(13, 'ADSFAsssssssssssss', 453, 'DSHFJTRSRW_lk;kjl');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `shoppinglist`
--
ALTER TABLE `shoppinglist`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `shoppinglist`
--
ALTER TABLE `shoppinglist`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
