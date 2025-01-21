import { createI18n } from 'vue-i18n';

// Импорт локалей
import en from '@/locales/en.json';
import ru from '@/locales/ru.json';
import zh from "@/locales/zh.json";

// Конфигурация i18n
const i18n = createI18n({
    locale: 'ru', // Язык по умолчанию
    fallbackLocale: 'zh', // Резервный язык
    messages: {
        en,
        ru,
        zh,
    },
});

export default i18n;
