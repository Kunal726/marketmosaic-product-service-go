-- MySQL dump 10.13  Distrib 9.2.0, for macos13.7 (arm64)
--
-- Host: localhost    Database: MarketMosaic
-- ------------------------------------------------------
-- Server version	9.2.0

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
-- Table structure for table `marketmosaic_admins`
--

DROP TABLE IF EXISTS `marketmosaic_admins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_admins` (
  `user_id` bigint NOT NULL,
  `account_verified` bit(1) DEFAULT NULL,
  `admin_notes` varchar(255) DEFAULT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  `last_login` datetime(6) DEFAULT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `FKrixew4pkb36ryq4b6ve7cpna1` FOREIGN KEY (`user_id`) REFERENCES `marketmosaic_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_admins`
--

LOCK TABLES `marketmosaic_admins` WRITE;
/*!40000 ALTER TABLE `marketmosaic_admins` DISABLE KEYS */;
/*!40000 ALTER TABLE `marketmosaic_admins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_category`
--

DROP TABLE IF EXISTS `marketmosaic_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_category` (
  `category_id` bigint NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL,
  `parent_id` bigint DEFAULT NULL,
  PRIMARY KEY (`category_id`),
  KEY `FKkhsykbdw5n0wyt1m80dn9dxh0` (`parent_id`),
  CONSTRAINT `FKkhsykbdw5n0wyt1m80dn9dxh0` FOREIGN KEY (`parent_id`) REFERENCES `marketmosaic_category` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_category`
--

LOCK TABLES `marketmosaic_category` WRITE;
/*!40000 ALTER TABLE `marketmosaic_category` DISABLE KEYS */;
INSERT INTO `marketmosaic_category` VALUES (1,'Electronics',NULL),(2,'Clothing',NULL),(3,'Home Goods',NULL),(4,'Toys',NULL),(5,'Books',NULL),(6,'Mobile Phones',1),(7,'Laptops',1),(8,'Accessories',1),(9,'Men\'s Wear',2),(10,'Women\'s Wear',2),(11,'Kids\' Wear',2),(12,'Furniture',3),(13,'Kitchenware',3),(14,'Action Figures',4),(15,'Educational Toys',4),(16,'Fiction',5),(17,'Non-Fiction',5),(18,'Android Phones',6),(19,'iPhones',6),(20,'Gaming Laptops',7),(21,'Business Laptops',7),(22,'Casual Wear',9),(23,'Formal Wear',9),(24,'Dresses',10),(25,'Tops',10);
/*!40000 ALTER TABLE `marketmosaic_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_product`
--

DROP TABLE IF EXISTS `marketmosaic_product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_product` (
  `product_id` bigint NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` decimal(38,2) NOT NULL,
  `stock_quantity` int DEFAULT '0',
  `supplier_id` bigint DEFAULT NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `date_added` datetime DEFAULT CURRENT_TIMESTAMP,
  `is_active` tinyint(1) DEFAULT '1',
  `category_id` bigint DEFAULT NULL,
  `rating` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`product_id`),
  KEY `SupplierID` (`supplier_id`),
  KEY `FKlb54eqxpxe3bkqu5f6465tljo` (`category_id`),
  CONSTRAINT `FKlb54eqxpxe3bkqu5f6465tljo` FOREIGN KEY (`category_id`) REFERENCES `marketmosaic_category` (`category_id`),
  CONSTRAINT `marketmosaic_product_ibfk_1` FOREIGN KEY (`supplier_id`) REFERENCES `marketmosaic_users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_product`
--

LOCK TABLES `marketmosaic_product` WRITE;
/*!40000 ALTER TABLE `marketmosaic_product` DISABLE KEYS */;
INSERT INTO `marketmosaic_product` VALUES (16,'OnePlus 9','OnePlus 9 with Snapdragon 888',893.99,120,5,'image_url_3','2025-02-09 14:23:38',1,1,''),(17,'Google Pixel 5','Google Pixel 5 with 5G',2993.99,45,5,'image_url_4','2025-02-09 14:23:38',1,1,''),(18,'Samsung Galaxy Note 20','Samsung Galaxy Note 20 Ultra',1539.99,90,4,'image_url_5','2025-02-09 14:23:38',1,1,''),(19,'Xiaomi Mi 11','Xiaomi Mi 11 with AMOLED Display',1739.99,55,5,'image_url_6','2025-02-09 14:23:38',1,1,''),(20,'Dell XPS 13','Premium Laptop',1329.99,85,4,'image_url_7','2025-02-09 14:23:38',1,2,''),(23,'Lenovo ThinkPad X1','Lenovo ThinkPad X1 Carbon',1999.99,35,5,'image_url_10','2025-02-09 14:23:38',1,2,NULL),(25,'Asus ZenBook 14','Asus ZenBook 14 with Intel Core i7',19569.99,60,5,'image_url_12','2025-02-09 14:23:38',1,2,''),(26,'Men\'s Casual T-Shirt','Comfortable Cotton T-Shirt',19.99,100,4,'image_url_13','2025-02-09 14:23:38',1,3,NULL),(27,'Men\'s Formal Shirt','Formal Shirt for Office Wear',39.99,60,5,'image_url_14','2025-02-09 14:23:38',1,3,NULL),(28,'Men\'s Jeans','Comfortable Denim Jeans',49.99,80,4,'image_url_15','2025-02-09 14:23:38',1,3,NULL),(29,'Men\'s Hoodie','Cozy Hoodie for Winter',59.99,50,5,'image_url_16','2025-02-09 14:23:38',1,3,NULL),(30,'Men\'s Sneakers','Stylish Men\'s Sneakers',89.99,70,4,'image_url_17','2025-02-09 14:23:38',1,3,NULL),(31,'Men\'s Jacket','Winter Jacket for Cold Weather',129.99,40,5,'image_url_18','2025-02-09 14:23:38',1,3,NULL),(32,'Women\'s Elegant Dress','Stylish Summer Dress',49.99,80,4,'image_url_19','2025-02-09 14:23:38',1,4,NULL),(33,'Women\'s Casual Top','Casual Top for Everyday Wear',29.99,100,5,'image_url_20','2025-02-09 14:23:38',1,6,NULL),(34,'Women\'s Skirt','Fashionable Skirt for Parties',39.99,60,4,'image_url_21','2025-02-09 14:23:38',1,4,NULL),(35,'Women\'s Sweater','Cozy Sweater for Cold Weather',49.99,40,5,'image_url_22','2025-02-09 14:23:38',1,4,NULL),(36,'Women\'s Jeans','Stylish Denim Jeans for Women',59.99,50,4,'image_url_23','2025-02-09 14:23:38',1,4,NULL),(37,'Women\'s High Heels','Elegant High Heels for Formal Events',89.99,30,5,'image_url_24','2025-02-09 14:23:38',1,4,NULL),(38,'The Silent Patient','Psychological Thriller Book',14.99,200,4,'image_url_25','2025-02-09 14:23:38',1,5,NULL),(39,'Atomic Habits','A Book on Building Better Habits',19.99,150,5,'image_url_26','2025-02-09 14:23:38',1,5,NULL),(40,'The Subtle Art of Not Giving a F*ck','Self-help Book on Living a Better Life',18.99,180,4,'image_url_27','2025-02-09 14:23:38',1,5,NULL),(41,'Becoming','Michelle Obama\'s Memoir',24.99,120,5,'image_url_28','2025-02-09 14:23:38',1,5,NULL),(42,'Educated','A Memoir by Tara Westover',22.99,140,4,'image_url_29','2025-02-09 14:23:38',1,19,NULL),(43,'The Alchemist','A Philosophical Novel by Paulo Coelho',16.99,160,5,'image_url_30','2025-02-09 14:23:38',1,5,NULL),(44,'Samsung Galaxy S21','The latest flagship phone from Samsung',799.99,50,4,NULL,'2025-02-09 14:34:01',1,1,NULL),(45,'Samsung Galaxy S21','The latest flagship phone from Samsung',799.99,50,4,NULL,'2025-02-09 14:35:34',1,19,NULL),(46,'iPhone 12','Apple\'s flagship phone',899.99,30,4,NULL,'2025-02-09 14:35:34',1,1,NULL),(47,'Men\'s Casual T-Shirt','Comfortable cotton t-shirt for men',19.99,100,5,NULL,'2025-02-09 14:35:34',1,2,NULL),(48,'Samsung Galaxy S21','The latest flagship phone from Samsung',799.99,50,4,NULL,'2025-02-09 14:35:41',1,20,NULL),(49,'iPhone 12','Apple\'s flagship phone',899.99,30,4,NULL,'2025-02-09 14:35:41',1,1,NULL),(50,'Men\'s Casual T-Shirt','Comfortable cotton t-shirt for men',19.99,100,5,NULL,'2025-02-09 14:35:41',1,2,NULL),(51,'Samsung Galaxy S21','The latest flagship phone from Samsung, featuring a 6.2-inch display, Snapdragon 888 chipset, and 5G support.',799.99,100,4,NULL,'2025-02-09 14:52:52',1,1,NULL),(52,'Samsung Galaxy S31','The latest flagship phone from Samsung, featuring a 6.2-inch display, Snapdragon 888 chipset, and 5G support.',799.99,100,4,NULL,'2025-02-09 14:54:35',1,1,NULL),(53,'Men\'s Casual T-Shirt','Comfortable and stylish casual t-shirt perfect for everyday wear.',19.99,120,4,'https://example.com/images/mens-casual-t-shirt.jpg','2025-02-09 14:56:20',1,2,NULL),(54,'Men\'s Formal Shirt','A classic formal shirt made from high-quality cotton.',39.99,80,4,'https://example.com/images/mens-formal-shirt.jpg','2025-02-09 14:56:20',1,2,NULL),(55,'Men\'s Jeans','Stylish and comfortable denim jeans for casual or semi-formal occasions.',49.99,150,5,'https://example.com/images/mens-jeans.jpg','2025-02-09 14:56:20',1,2,NULL),(56,'Men\'s Hoodie','Cozy and warm hoodie, perfect for chilly days or casual outings.',29.99,100,5,'https://example.com/images/mens-hoodie.jpg','2025-02-09 14:56:20',1,2,NULL),(57,'test','hgdh',50.00,20,4,NULL,'2025-02-10 10:21:38',1,2,NULL),(58,'test','hgdh',50.00,20,4,NULL,'2025-02-10 10:23:17',1,2,NULL),(59,'xew','Striexwng',60.00,0,4,'String','2025-02-10 10:43:38',1,6,NULL),(60,'xee2','ezdxe',21.00,0,4,'x  ','2025-02-27 16:00:19',0,3,NULL),(61,'','ezdxe',21.00,0,4,'x  ','2025-02-27 16:10:46',0,3,NULL),(62,'dq33fr','',21.00,0,4,'x  ','2025-02-27 16:12:01',0,3,NULL),(63,'rx3gc',NULL,13.00,0,4,'x  ','2025-02-27 17:36:16',0,3,NULL),(64,'rx3gc',NULL,13.00,0,4,'x  ','2025-02-27 17:38:34',0,3,NULL),(65,'rx3gc',NULL,13.00,0,4,'x  ','2025-02-27 17:40:18',0,3,NULL),(66,'rx3gc',NULL,13.00,0,4,'x  ','2025-02-27 17:43:13',0,3,NULL),(67,'rx3gc',NULL,13.00,0,4,'x  ','2025-02-27 17:47:22',0,3,NULL),(68,'ewqr  ',NULL,13.00,0,4,'x  ','2025-02-28 12:20:36',0,3,NULL),(69,'ewqr  ',NULL,13.00,0,4,'x  ','2025-02-28 12:22:42',0,3,NULL),(70,'ewqr n ',NULL,13.00,0,4,'x  ','2025-02-28 12:23:40',0,3,NULL),(71,'Premium Headphones','High-quality wireless headphones with noise cancellation',199.99,50,4,'https://example.com/images/headphones.jpg','2025-03-13 11:29:58',1,1,''),(72,'Premium Headphones','High-quality wireless headphones with noise cancellation',199.99,50,4,'https://example.com/images/headphones.jpg','2025-03-13 11:33:55',1,1,''),(73,'Premium h','High-quality wireless headphones with noise cancellation',199.99,50,4,'https://example.com/images/headphones.jpg','2025-03-13 11:35:03',1,1,''),(74,'Premium Headphones Pro','High-quality wireless headphones with noise cancellation and 40-hour battery life',199.99,50,4,'https://example.com/images/headphones-pro.jpg','2025-03-13 13:29:23',1,1,''),(75,'Wireless Earbuds Elite','True wireless earbuds with premium sound quality and touch controls',149.99,75,4,'https://example.com/images/earbuds-elite.jpg','2025-03-13 13:29:23',1,1,''),(76,'Gaming Headset X1','Professional gaming headset with 7.1 surround sound and RGB lighting',129.99,60,4,'https://example.com/images/gaming-headset.jpg','2025-03-13 13:29:23',1,1,''),(77,'Studio Monitor Headphones','Professional-grade studio monitoring headphones for audio production',249.99,40,4,'https://example.com/images/studio-headphones.jpg','2025-03-13 13:29:23',1,1,''),(78,'Sport Wireless Earphones','Sweat-resistant wireless earphones perfect for workouts and running',89.99,100,4,'https://example.com/images/sport-earphones.jpg','2025-03-13 13:29:23',1,1,''),(79,'Premium Headphones Pro','High-quality wireless headphones with noise cancellation and 40-hour battery life',199.99,50,4,'https://example.com/images/headphones-pro.jpg','2025-03-13 13:59:35',1,1,''),(80,'Wireless Earbuds Elite','True wireless earbuds with premium sound quality and touch controls',149.99,75,4,'https://example.com/images/earbuds-elite.jpg','2025-03-13 13:59:35',1,1,''),(81,'Gaming Headset X1','Professional gaming headset with 7.1 surround sound and RGB lighting',129.99,60,4,'https://example.com/images/gaming-headset.jpg','2025-03-13 13:59:35',1,1,''),(82,'Studio Monitor Headphones','Professional-grade studio monitoring headphones for audio production',249.99,40,4,'https://example.com/images/studio-headphones.jpg','2025-03-13 13:59:35',1,1,''),(83,'Sport Wireless Earphones','Sweat-resistant wireless earphones perfect for workouts and running',89.99,100,4,'https://example.com/images/sport-earphones.jpg','2025-03-13 13:59:35',1,22,''),(84,'Premium h','High-quality wireless headphones with noise cancellation',199.99,50,4,'https://example.com/images/headphones.jpg','2025-03-13 14:00:25',1,18,''),(85,'g','High-quality wireless headphones with noise cancellation',199.99,50,4,'https://example.com/images/headphones.jpg','2025-03-13 16:07:10',1,18,''),(86,'g','High-quality wireless headphones with noise cancellation',21.00,50,4,'https://example.com/images/headphones.jpg','2025-03-13 16:20:20',1,18,''),(87,'dxg','High-quality wireless headphones with noise cancellation',21.00,50,4,'https://example.com/images/headphones.jpg','2025-03-13 16:36:06',1,18,''),(88,'Premium Headphones Pro','High-quality wireless headphones with noise cancellation and 40-hour battery life',199.99,50,4,'https://example.com/images/headphones-pro.jpg','2025-03-13 16:36:16',1,1,''),(89,'Wireless Earbuds Elite','True wireless earbuds with premium sound quality and touch controls',149.99,75,4,'https://example.com/images/earbuds-elite.jpg','2025-03-13 16:36:16',1,1,''),(90,'Gaming Headset X1','Professional gaming headset with 7.1 surround sound and RGB lighting',129.99,60,4,'https://example.com/images/gaming-headset.jpg','2025-03-13 16:36:16',1,1,''),(91,'Studio Monitor Headphones','Professional-grade studio monitoring headphones for audio production',249.99,40,4,'https://example.com/images/studio-headphones.jpg','2025-03-13 16:36:16',1,1,''),(92,'Sport Wireless Earphones','Sweat-resistant wireless earphones perfect for workouts and running',89.99,100,4,'https://example.com/images/sport-earphones.jpg','2025-03-13 16:36:16',1,22,''),(93,'dxg','High-quality wireless headphones with noise cancellation',21.00,50,4,'https://example.com/images/headphones.jpg','2025-03-13 17:17:33',1,18,''),(94,'Premium Headphones Pro','High-quality wireless headphones with noise cancellation and 40-hour battery life',199.99,50,4,'https://example.com/images/headphones-pro.jpg','2025-03-13 17:18:01',1,1,''),(95,'Wireless Earbuds Elite','True wireless earbuds with premium sound quality and touch controls',149.99,75,4,'https://example.com/images/earbuds-elite.jpg','2025-03-13 17:18:01',1,1,''),(96,'Gaming Headset X1','Professional gaming headset with 7.1 surround sound and RGB lighting',129.99,60,4,'https://example.com/images/gaming-headset.jpg','2025-03-13 17:18:01',1,1,''),(97,'Studio Monitor Headphones','Professional-grade studio monitoring headphones for audio production',249.99,40,4,'https://example.com/images/studio-headphones.jpg','2025-03-13 17:18:01',1,1,''),(98,'Sport Wireless Earphones','Sweat-resistant wireless earphones perfect for workouts and running',89.99,100,4,'https://example.com/images/sport-earphones.jpg','2025-03-13 17:18:01',1,22,''),(99,'dxg','High-quality wireless headphones with noise cancellation',21.00,50,4,'https://example.com/images/headphones.jpg','2025-03-13 17:33:32',1,18,''),(100,'Premium Headphones Pro','High-quality wireless headphones with noise cancellation and 40-hour battery life',199.99,50,4,'https://example.com/images/headphones-pro.jpg','2025-03-13 17:33:48',1,1,''),(101,'Wireless Earbuds Elite','True wireless earbuds with premium sound quality and touch controls',149.99,75,4,'https://example.com/images/earbuds-elite.jpg','2025-03-13 17:33:48',1,1,''),(102,'Gaming Headset X1','Professional gaming headset with 7.1 surround sound and RGB lighting',129.99,60,4,'https://example.com/images/gaming-headset.jpg','2025-03-13 17:33:48',1,1,''),(103,'Studio Monitor Headphones','Professional-grade studio monitoring headphones for audio production',249.99,40,4,'https://example.com/images/studio-headphones.jpg','2025-03-13 17:33:48',1,1,''),(104,'Sport Wireless Earphones','Sweat-resistant wireless earphones perfect for workouts and running',89.99,100,4,'https://example.com/fimages/sport-earphones.jpg','2025-03-13 17:33:48',1,22,''),(105,'Sunny','High-quality wireless headphones with noise cancellation',21.00,50,4,'https://example.com/images/headphones.jpg','2025-04-20 16:54:52',1,18,''),(106,'Suneny','High-quality wireless headphones with noise cancellation',21.00,50,4,'https://example.com/images/headphones.jpg','2025-04-20 17:02:08',1,18,'');
/*!40000 ALTER TABLE `marketmosaic_product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_product_images`
--

DROP TABLE IF EXISTS `marketmosaic_product_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_product_images` (
  `image_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint NOT NULL,
  `url` varchar(255) NOT NULL,
  `alt_text` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`image_id`),
  KEY `idx_product_id` (`product_id`),
  CONSTRAINT `marketmosaic_product_images_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `marketmosaic_product` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_product_images`
--

LOCK TABLES `marketmosaic_product_images` WRITE;
/*!40000 ALTER TABLE `marketmosaic_product_images` DISABLE KEYS */;
/*!40000 ALTER TABLE `marketmosaic_product_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_product_videos`
--

DROP TABLE IF EXISTS `marketmosaic_product_videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_product_videos` (
  `video_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `product_id` bigint NOT NULL,
  `url` varchar(255) NOT NULL,
  `thumbnail_url` varchar(255) DEFAULT NULL,
  `file_name` varchar(255) NOT NULL,
  PRIMARY KEY (`video_id`),
  KEY `idx_product_id` (`product_id`),
  CONSTRAINT `marketmosaic_product_videos_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `marketmosaic_product` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_product_videos`
--

LOCK TABLES `marketmosaic_product_videos` WRITE;
/*!40000 ALTER TABLE `marketmosaic_product_videos` DISABLE KEYS */;
/*!40000 ALTER TABLE `marketmosaic_product_videos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_sellers`
--

DROP TABLE IF EXISTS `marketmosaic_sellers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_sellers` (
  `user_id` bigint NOT NULL,
  `address` varchar(255) DEFAULT NULL,
  `business_name` varchar(255) NOT NULL,
  `contact_phone` varchar(255) DEFAULT NULL,
  `seller_status` enum('APPROVED','NONE','PENDING','REJECTED','REVOKED') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `FK176vdler6gqu3su0brmm1qpk9` FOREIGN KEY (`user_id`) REFERENCES `marketmosaic_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_sellers`
--

LOCK TABLES `marketmosaic_sellers` WRITE;
/*!40000 ALTER TABLE `marketmosaic_sellers` DISABLE KEYS */;
/*!40000 ALTER TABLE `marketmosaic_sellers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_shipping_addresses`
--

DROP TABLE IF EXISTS `marketmosaic_shipping_addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_shipping_addresses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `address_line1` varchar(255) NOT NULL,
  `address_line2` varchar(255) DEFAULT NULL,
  `address_type` varchar(255) DEFAULT NULL,
  `city` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  `is_active` bit(1) NOT NULL,
  `is_default` bit(1) NOT NULL,
  `phone_number` varchar(255) NOT NULL,
  `postal_code` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FKcc8ry22oxx0750j8sj7yt1llv` (`user_id`),
  CONSTRAINT `FKcc8ry22oxx0750j8sj7yt1llv` FOREIGN KEY (`user_id`) REFERENCES `marketmosaic_users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_shipping_addresses`
--

LOCK TABLES `marketmosaic_shipping_addresses` WRITE;
/*!40000 ALTER TABLE `marketmosaic_shipping_addresses` DISABLE KEYS */;
/*!40000 ALTER TABLE `marketmosaic_shipping_addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_tag`
--

DROP TABLE IF EXISTS `marketmosaic_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_tag` (
  `tag_id` bigint NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) NOT NULL,
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_tag`
--

LOCK TABLES `marketmosaic_tag` WRITE;
/*!40000 ALTER TABLE `marketmosaic_tag` DISABLE KEYS */;
INSERT INTO `marketmosaic_tag` VALUES (1,'New Arrival'),(2,'Best Seller'),(3,'Discount'),(4,'Featured');
/*!40000 ALTER TABLE `marketmosaic_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `marketmosaic_users`
--

DROP TABLE IF EXISTS `marketmosaic_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `marketmosaic_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `roles` varchar(255) DEFAULT 'ROLE_USER',
  `date_added` datetime DEFAULT CURRENT_TIMESTAMP,
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` datetime(6) DEFAULT NULL,
  `last_login` datetime(6) DEFAULT NULL,
  `phone_number` varchar(255) DEFAULT NULL,
  `profile_picture` varchar(255) DEFAULT NULL,
  `shipping_address` varchar(255) DEFAULT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UKdxpo4thvwoqhwcpncq5h7twty` (`email`),
  UNIQUE KEY `UK4vhrtlr9cjtsrr7nph1go299x` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `marketmosaic_users`
--

LOCK TABLES `marketmosaic_users` WRITE;
/*!40000 ALTER TABLE `marketmosaic_users` DISABLE KEYS */;
INSERT INTO `marketmosaic_users` VALUES (1,'abc@gmail.com','test','test123','test','ROLE_USER','2025-02-08 11:18:43',1,NULL,NULL,NULL,NULL,NULL,NULL),(3,'abfc@gmail.com','test1','test123','test1','ROLE_USER','2025-02-08 11:18:43',1,NULL,NULL,NULL,NULL,NULL,NULL),(4,'testsun@gmail.com','tersy','$2a$12$7Oe6d6.c2SF34R/VUXBmoue6jXNwsOFgZpbrRIG5EgyCQTUjuRoD6','test1234','ROLE_USER','2025-02-08 11:18:43',1,NULL,NULL,NULL,NULL,NULL,NULL),(5,'Sunny','Sunny','$2a$10$HiZ1ZdZ5Tt4nexb56oPPq.K7jBhLaGIi5kANDn2WcB2HczSgUubia','Sunny','ROLE_USER','2025-02-08 11:18:43',1,NULL,NULL,NULL,NULL,NULL,NULL),(6,'SunnTech','Sunny Sonar','$2a$10$QGfjPYBMzy85OOYKUNZdOu88Noinz/gT2x8J202q/NpPeuKjVpCzu','SunnTech','ROLE_USER','2025-03-14 09:33:15',1,NULL,NULL,NULL,NULL,NULL,NULL),(7,'admin','Sunny Sonar','$2a$10$zPBs9qvm.zkSNAFQzkqwp.x9XL72lwuv87b7sS1Bm8ktVGxMIxCES','admin','ROLE_ADMIN','2025-03-16 10:36:31',1,NULL,NULL,NULL,NULL,NULL,NULL),(9,'techsunny121@gmail.com','Sunny','$2a$10$yCuioId9Fw4paWYXFOxY1OtQqzqikvjYk.Go57KrxFGbO9KRcEIBC','Sun','ROLE_ADMIN','2025-03-29 18:14:15',1,'2025-03-29 18:14:14.940606','2025-04-22 12:11:16.400441',NULL,'Sun_47dd2c05-ab30-4940-80a5-586209f25b80.jpeg',NULL,'2025-04-22 12:11:16.407603');
/*!40000 ALTER TABLE `marketmosaic_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `password_reset_tokens`
--

DROP TABLE IF EXISTS `password_reset_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_reset_tokens` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `expiry_date` datetime(6) NOT NULL,
  `token` varchar(255) NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK71lqwbwtklmljk3qlsugr1mig` (`token`),
  KEY `FKae6nku0qg2ip2wgcbmwyi2n56` (`user_id`),
  CONSTRAINT `FKae6nku0qg2ip2wgcbmwyi2n56` FOREIGN KEY (`user_id`) REFERENCES `marketmosaic_users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `password_reset_tokens`
--

LOCK TABLES `password_reset_tokens` WRITE;
/*!40000 ALTER TABLE `password_reset_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `password_reset_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_tags`
--

DROP TABLE IF EXISTS `product_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_tags` (
  `product_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL,
  KEY `FK3bkb2xuovq56nk7xka5aq0wkb` (`tag_id`),
  KEY `FKqgurxrtw8vpi4n4nf2p1hax2p` (`product_id`),
  CONSTRAINT `FK3bkb2xuovq56nk7xka5aq0wkb` FOREIGN KEY (`tag_id`) REFERENCES `marketmosaic_tag` (`tag_id`),
  CONSTRAINT `FKqgurxrtw8vpi4n4nf2p1hax2p` FOREIGN KEY (`product_id`) REFERENCES `marketmosaic_product` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_tags`
--

LOCK TABLES `product_tags` WRITE;
/*!40000 ALTER TABLE `product_tags` DISABLE KEYS */;
INSERT INTO `product_tags` VALUES (16,1),(17,2),(18,3),(19,1),(20,2),(23,2),(25,2),(26,1),(27,4),(28,3),(29,1),(30,2),(31,4),(32,1),(33,2),(34,3),(35,1),(36,2),(37,4),(38,2),(39,1),(40,2),(41,4),(42,3),(43,2),(51,1),(51,2),(53,1),(53,2),(53,3),(53,4),(54,1),(54,2),(54,4),(55,1),(55,3),(55,4),(56,2),(56,3),(56,4),(57,1),(57,2),(58,1),(58,2),(72,1),(72,2),(72,3),(73,1),(73,2),(73,3),(74,1),(74,2),(74,3),(75,1),(75,2),(76,1),(76,3),(77,2),(77,3),(78,1),(79,1),(79,2),(79,3),(80,1),(80,2),(81,1),(81,3),(82,2),(82,3),(83,1),(84,1),(84,2),(84,3),(85,1),(85,2),(85,3),(86,1),(86,2),(86,3),(87,1),(87,2),(87,3),(88,1),(88,2),(88,3),(89,1),(89,2),(90,1),(90,3),(91,2),(91,3),(92,1),(93,1),(93,2),(93,3),(94,1),(94,2),(94,3),(95,1),(95,2),(96,1),(96,3),(97,2),(97,3),(98,1),(99,1),(99,2),(99,3),(100,1),(100,2),(100,3),(101,1),(101,2),(102,1),(102,3),(103,2),(103,3),(104,1),(105,1),(105,2),(105,3),(106,1),(106,2),(106,3);
/*!40000 ALTER TABLE `product_tags` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-11 19:59:50
