
# [Gym Log](https://workout.lkhtk.me/)

**[Gym Log](https://workout.lkhtk.me/)** — это полнофункциональное веб-приложение для отслеживания тренировок. Оно позволяет пользователям записывать, просматривать и анализировать свои тренировки, обеспечивая удобный интерфейс и надежную серверную часть.

---

## 🧰 Используемые технологии

### 🔙 Backend

- **Go (Golang)** — основной язык серверной логики.
- **Gin** — высокопроизводительный HTTP-фреймворк.
- **MongoDB** — база данных.
- **Docker** — контейнеризация приложения.

### 🔜 Frontend

- **Vue.js** — JavaScript-фреймворк для пользовательского интерфейса.
- **Pinia** — управление состоянием.
- **Vue Router** — маршрутизация на клиенте.
- **Axios** — HTTP-клиент.
- **Docker** — для контейнеризации клиентской части.

### ⚙️ DevOps

- **Docker Compose** — оркестрация контейнеров.
- **GitHub Actions** — автоматизация CI/CD.
- **Helm** — деплой в Kubernetes.
- **Kubernetes** — управление кластером.

---

## 🚀 Установка и запуск

### 📦 Зависимости

Перед запуском убедитесь, что у вас установлены:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/)
- [Node.js и npm](https://nodejs.org/)

---

### 🔧 Быстрый старт с Docker Compose

```bash
git clone https://github.com/lkhtk/workout-log.git
cd workout-log
docker-compose up --build
```

После сборки приложение будет доступно по адресу:

```
http://localhost:8080
```

---

### 🛠️ Ручная сборка

#### Backend (Go)

```bash
cd backend
go mod tidy
go run main.go
```

#### Frontend (Vue.js)

```bash
cd frontend
npm install
npm run serve
```

Доступ: `http://localhost:8080`

---

### ☁️ Деплой в Kubernetes с Helm

1. Установите Helm: https://helm.sh/docs/intro/install/
2. Настройте доступ к вашему кластеру Kubernetes.
3. Установите приложение:

```bash
cd chart/workout-log
helm install workout-log .
```

---

## 🗂 Структура проекта

```
workout-log/
├── backend/               # Серверная часть на Go
├── frontend/              # Клиентская часть на Vue.js
├── chart/workout-log/     # Helm-чарт для Kubernetes
├── docker-compose.yml     # Конфигурация Docker Compose
└── README.md              # Документация проекта
```

---

## 📄 Лицензия

Этот проект лицензирован под MIT License. Подробности читайте в файле [LICENSE](https://github.com/lkhtk/workout-log/blob/main/LICENSE).

---

## 📬 Обратная связь

Для багов, предложений или pull request'ов — заходите в [репозиторий на GitHub](https://github.com/lkhtk/workout-log).
