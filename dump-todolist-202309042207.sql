-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: todolist
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `todo_items`
--

DROP TABLE IF EXISTS `todo_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `image` json DEFAULT NULL,
  `description` text,
  `status` enum('Doing','Done','Deleted') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'Doing',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `todo_items`
--

LOCK TABLES `todo_items` WRITE;
/*!40000 ALTER TABLE `todo_items` DISABLE KEYS */;
INSERT INTO `todo_items` VALUES (1,'This is task 1',NULL,'This is task 1 description, abcd','Doing','2023-08-31 08:43:34','2023-08-31 08:43:34'),(2,'This is task 2',NULL,'This is task 2 description, abcd','Done','2023-08-31 08:44:00','2023-09-03 09:57:43'),(3,'This is task 3',NULL,'This is task 3 description, abcd','Deleted','2023-08-31 08:44:00','2023-09-03 09:57:43'),(4,'This is task 4',NULL,'This is task 4 description, abcd','Doing','2023-08-31 08:44:00','2023-08-31 08:44:00'),(5,'This is task 5',NULL,'This is task 5 description, abcd','Doing','2023-08-31 08:44:00','2023-08-31 08:44:00'),(6,'This is task 6',NULL,'This is task 6 description, abcd222','Doing','2023-08-31 08:44:00','2023-08-31 08:47:02'),(7,'This is a new item',NULL,'item description','Doing','2023-09-03 08:28:00','2023-09-03 08:28:00'),(8,'This is an update 2 item 2',NULL,'update item description 2','Deleted','2023-09-03 08:30:55','2023-09-03 14:39:44'),(10,'This is an update 2 item 10',NULL,'update item description 10','Done','2023-09-03 17:03:21','2023-09-03 17:03:55'),(11,'This is a new item 11',NULL,'item description 11','Doing','2023-09-04 14:31:26','2023-09-04 14:31:26'),(12,'This is a new item 12',NULL,'item description 12','Doing','2023-09-04 14:35:05','2023-09-04 14:35:05'),(13,'This is a new item 13',NULL,'item description 13','Done','2023-09-04 14:36:46','2023-09-04 14:36:46');
/*!40000 ALTER TABLE `todo_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'todolist'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-04 22:07:13
