# 🥗 Yummi-FE

**Yummi-FE** is a beautifully designed React Native app that helps users **track meals, analyze nutrition**, and **promote healthy eating habits**. Built using [Expo](https://expo.dev/), Tailwind CSS (via NativeWind), and Framer Motion, it offers a smooth and intuitive experience. The app fetches recipes from [TheMealDB API](https://www.themealdb.com/api/json/v1/1) and communicates with a custom backend (`Yummi-BE`) written in Go and deployed on AWS.

> Im going to published this app on **F-Droid** and the **Galaxy Store** for maximum reach and accessibility.

---

## 📱 Features

* 🍛 Search and explore meals using [TheMealDB API](https://www.themealdb.com/api.php)
* 📊 Nutrition insights and meal breakdowns
* 💚 Calorie & macro tracking (Planned)
* 🔐 Secure authentication via [Clerk](https://clerk.dev/)
* 🌐 RESTful integration with a Go-powered backend hosted on AWS
* 💅 Smooth animations and sleek UI with Framer Motion
* 📦 Fully mobile-first and optimized for Android deployment

---

## 🧹 Tech Stack

| Layer      | Technology                                           |
| ---------- | ---------------------------------------------------- |
| Language   | React Native (TypeScript)                            |
| Framework  | Expo + Expo Router                                   |
| UI         | Tailwind CSS (NativeWind) + Framer Motion            |
| Backend    | Go (`Yummi-BE`)                                      |
| API        | [TheMealDB](https://www.themealdb.com/api/json/1) |
| Auth       | Clerk for secure login                               |
| Deployment | AWS (Backend) + F-Droid/Galaxy Store (Upcoming)      |

---

## 📂 Project Structure

```
mobile/
├── app/                # Expo router pages
├── components/         # Reusable UI components
├── constants/          # API routes & global configs
├── hooks/              # Custom React hooks
├── assets/             # Fonts, images, icons
├── scripts/            # Utility scripts (e.g., reset-project)
├── .env                # Environment variables (ignored in Git)
└── tailwind.config.js  # TailwindCSS config
```

---

## 🛠️ Installation

```bash
# Clone the repo
git clone https://github.com/GitNinja36/Yummi-FE.git
cd Yummi-FE/mobile

# Install dependencies
npm install

# Start development server
npm run android    # For Android
npm run ios        # For iOS (Mac only)
npm run web        # For web preview
```

---

## 📸 App Icon & Store Readiness

The app includes a custom logo and splash screen for store listings. Make sure your `app.json` includes:

```json
{
  "expo": {
    "name": "Yummi",
    "slug": "yummi",
    "icon": "./assets/icon.png",
    "splash": {
      "image": "./assets/splash.png",
      "resizeMode": "contain",
      "backgroundColor": "#ffffff"
    }
  }
}
```

---

## 🚀 Upcoming Deployment

### 🌟 Target Stores:

* [x] F-Droid (open-source, free Android marketplace)
* [x] Galaxy Store

### 📟 Requirements Checklist:

* [ ] Signed APK or AAB
* [ ] Unique app ID (`com.gitninja36.yummi`)
* [ ] Screenshots (1080x1920)
* [ ] App icon and splash screen
* [ ] Privacy policy and permissions declaration

---

## 📄 License

This project is licensed under the **MIT License**.

---

## 👨‍💻 Author

Developed with 💚 by [Rohit Kumar (GitNinja36)](https://github.com/GitNinja36)

---

## ⭐️ Support

If you find this project useful, consider giving it a ⭐️ on GitHub!

