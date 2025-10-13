📦 PriceAnalysisGoBot — Telegram бот для отслеживания цен

PriceAnalysisGoBot — это Telegram-бот, написанный на Go, который помогает находить и отслеживать цены товаров на маркетплейсах (Ozon, Wildberries, Kaspi и др.).
Вы можете искать товары, сравнивать цены и подписываться на уведомления о снижении стоимости.

🚀 Возможности

Поиск товара по названию:

`/price iPhone 14`

Ответ:

`iPhone 14:
• Ozon: 380 000₸
• Wildberries: 370 500₸
• Kaspi: 385 000₸`

Подписка на уведомления о снижении цены:

`/track iPhone 14 360000`

Список отслеживаемых товаров:

`/list`

Удаление подписки:

`/remove 1`

🛠 Технологии

* Go 1.22+
* PostgreSQL (через Docker)
* HTML-парсинг: `github.com/PuerkitoBio/goquery`
* Планировщик задач: `github.com/robfig/cron/v3`
* Работа с .env файлом: `github.com/joho/godotenv`