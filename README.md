# GO API

Para executar a API executar o comando: `go run main.go`
A API escuta a porta `3000`, para acessar utilize o endereço: `http://localhost:3000`

#### O que tem?

CRUDS:

- Users
- Sellers

#### Endpoint

##### Home
- [GET] /

##### Users
- [GET] /v1/users
- [GET] /v1/users/:id
- [POST] /v1/users
- [PUT] /v1/users/:id
- [DELETE] /v1/users/:id

##### Sellers
- [GET] /v1/sellers
- [GET] /v1/sellers/:id
- [POST] /v1/sellers
- [PUT] /v1/sellers/:id
- [DELETE] /v1/sellers/:id

#### BANCO DE DADOS

Para alterar os dados de conexão com o banco de dados altere o arquivo: `lib/db.go`

##### DUMP - MYSQL

```
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Copiando estrutura do banco de dados para goapi
DROP DATABASE IF EXISTS `goapi`;
CREATE DATABASE IF NOT EXISTS `goapi` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `goapi`;

-- Copiando estrutura para tabela goapi.sellers
DROP TABLE IF EXISTS `sellers`;
CREATE TABLE IF NOT EXISTS `sellers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- Exportação de dados foi desmarcado.
-- Copiando estrutura para tabela goapi.users
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- Exportação de dados foi desmarcado.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
```