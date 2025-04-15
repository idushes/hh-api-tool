# HeadHunter API Tools

Этот проект содержит инструменты для работы с API HeadHunter через MCP (Mission Control Platform).

## Планируемые tools MCP (TODO)

### Авторизация

- [ ] `POST /oauth/token` — Получение access-токена (авторизация)
- [ ] `DELETE /oauth/token` — Инвалидация токена

### Резюме

- [ ] `GET /resumes/mine` — Список резюме авторизованного пользователя
- [ ] `POST /resumes` — Создание резюме
- [ ] `GET /resumes/{resume_id}` — Просмотр резюме
- [ ] `PUT /resumes/{resume_id}` — Редактирование резюме
- [ ] `POST /resumes/{resume_id}/publish` — Публикация резюме
- [ ] `DELETE /resumes/{resume_id}` — Удаление резюме
- [ ] `GET /resumes/{resume_id}/similar_vacancies` — Похожие вакансии для резюме
- [ ] `GET /resumes/{resume_id}/views` — Просмотр статистики просмотров резюме
- [ ] `GET /resumes/creation-availability` — Проверка возможности создания резюме
- [ ] `GET /resumes/{vacancy_id}/suitable_resumes` — Список подходящих для отклика резюме
- [ ] `GET /resumes/{resume_id}/status` — Получение статуса резюме
- [ ] `GET /resumes/{resume_id}/conditions` — Условия заполнения полей резюме
- [ ] `GET /resumes/{resume_id}/negotiations_history` — История откликов/приглашений по резюме

### Работа с телефоном в резюме

- [ ] `POST /resume_phone_request` — Запросить звонок с кодом подтверждения
- [ ] `POST /resume_phone_confirm` — Подтвердить телефон кодом

### Поиск резюме (для работодателей)

- [ ] `GET /resumes` — Поиск резюме
- [ ] `POST /saved_searches/resumes` — Сохранение поиска резюме
- [ ] `GET /saved_searches/resumes` — Список сохраненных поисков резюме
- [ ] `GET /saved_searches/resumes/{id}` — Получение единичного сохраненного поиска резюме
- [ ] `PUT /saved_searches/resumes/{id}` — Редактирование сохраненного поиска резюме
- [ ] `DELETE /saved_searches/resumes/{id}` — Удаление сохраненного поиска резюме
- [ ] `PUT /saved_searches/resumes/{saved_search_id}/managers/{manager_id}` — Передача сохраненного поиска резюме другому менеджеру

### Вакансии

- [ ] `POST /vacancies` — Публикация вакансии
- [ ] `GET /employers/{employer_id}/vacancies/active` — Просмотр списка опубликованных вакансий
- [ ] `GET /vacancies/{vacancy_id}` — Просмотр вакансии
- [ ] `PUT /vacancies/{vacancy_id}` — Редактирование вакансии
- [ ] `DELETE /vacancies/{vacancy_id}` — Удаление вакансии
- [ ] `GET /vacancies/{vacancy_id}/similar_vacancies` — Поиск похожих вакансий
- [ ] `GET /vacancies/{vacancy_id}/stats` — Просмотр статистики по вакансии
- [ ] `GET /employers/{employer_id}/vacancies/archived` — Список архивных вакансий
- [ ] `POST /employers/{employer_id}/vacancies/{vacancy_id}/archive` — Добавление вакансии в архив
- [ ] `POST /employers/{employer_id}/vacancies/{vacancy_id}/restore` — Восстановление вакансии из архива
- [ ] `GET /vacancy_conditions` — Условия заполнения полей вакансии

### Черновики вакансий

- [ ] `POST /vacancies/drafts` — Создание черновика вакансии
- [ ] `GET /vacancies/drafts/{draft_id}` — Получение черновика вакансии
- [ ] `PUT /vacancies/drafts/{draft_id}` — Обновление черновика вакансии
- [ ] `DELETE /vacancies/drafts/{draft_id}` — Удаление черновика вакансии
- [ ] `POST /vacancies/drafts/{draft_id}/publish` — Публикация вакансии из черновика
- [ ] `GET /vacancies/drafts/{draft_id}/duplicates` — Поиск дубликатов вакансии при публикации

### Поиск вакансий

- [ ] `GET /vacancies` — Поиск вакансий
- [ ] `POST /saved_searches/vacancies` — Сохранение поиска вакансий
- [ ] `GET /saved_searches/vacancies` — Список сохраненных поисков вакансий
- [ ] `GET /saved_searches/vacancies/{id}` — Получение единичного сохраненного поиска вакансий
- [ ] `PUT /saved_searches/vacancies/{id}` — Редактирование сохраненного поиска вакансий
- [ ] `DELETE /saved_searches/vacancies/{id}` — Удаление сохраненного поиска вакансий

### Переписка и отклики

- [ ] `GET /negotiations` — Список откликов/приглашений
- [ ] `POST /negotiations/{vacancy_id}/markup` — Отклик на вакансию с сопроводительным письмом
- [ ] `POST /negotiations/{negotiations_id}/messages` — Отправка сообщения в отклике/приглашении
- [ ] `GET /negotiations/{negotiation_id}/messages` — Просмотр списка сообщений в отклике/приглашении
- [ ] `PUT /negotiations/topics/{nid}` — Изменение статуса отклика
- [ ] `POST /negotiations/topics/{nid}/messages` — Отправка сообщения работодателю
- [ ] `GET /negotiations/topics/{nid}/messages` — Получение списка сообщений в отклике
- [ ] `DELETE /negotiations/topics/{nid}` — Скрытие отклика у соискателя
- [ ] `POST /negotiations/phone_calls/{phone_call_id}/decline` — Отклонение звонка
- [ ] `GET /negotiations/phone_calls/{id}` — Информация о звонке

### Работодатели

- [ ] `GET /employers/{employer_id}` — Информация о работодателе
- [ ] `GET /employers` — Поиск работодателей
- [ ] `GET /employers/{employer_id}/managers` — Просмотр списка менеджеров работодателя
- [ ] `GET /manager_accounts/mine` — Рабочие аккаунты менеджера
- [ ] `GET /employers/{employer_id}/manager_types` — Типы менеджеров работодателя
- [ ] `GET /employers/{employer_id}/departments` — Список отделов работодателя
- [ ] `GET /employers/{employer_id}/vacancy_branded_templates` — Список брендированных шаблонов вакансий

### Адреса работодателей

- [ ] `GET /employers/{employer_id}/addresses` — Список адресов работодателя
- [ ] `POST /employers/{employer_id}/addresses` — Добавление адреса работодателя
- [ ] `GET /employers/{employer_id}/addresses/{address_id}` — Получение адреса работодателя
- [ ] `PUT /employers/{employer_id}/addresses/{address_id}` — Редактирование адреса работодателя
- [ ] `DELETE /employers/{employer_id}/addresses/{address_id}` — Удаление адреса работодателя

### Избранное и скрытые элементы

- [ ] `GET /vacancies/favorited` — Список избранных вакансий соискателя
- [ ] `PUT /vacancies/favorited/{vacancy_id}` — Добавление вакансии в избранное
- [ ] `DELETE /vacancies/favorited/{vacancy_id}` — Удаление вакансии из избранного
- [ ] `GET /vacancies/blacklisted` — Список скрытых вакансий
- [ ] `PUT /vacancies/blacklisted/{vacancy_id}` — Добавление вакансии в список скрытых
- [ ] `DELETE /vacancies/blacklisted/{vacancy_id}` — Удаление вакансии из списка скрытых
- [ ] `GET /employers/blacklisted` — Список скрытых работодателей
- [ ] `PUT /employers/blacklisted/{employer_id}` — Добавление работодателя в список скрытых
- [ ] `DELETE /employers/blacklisted/{employer_id}` — Удаление работодателя из списка скрытых

### Артефакты (фото и портфолио)

- [ ] `POST /artifacts` — Загрузка артефакта
- [ ] `GET /artifacts/portfolio` — Получение портфолио
- [ ] `GET /artifacts/photo` — Получение фотографий
- [ ] `PUT /artifacts/{id}` — Редактирование артефакта
- [ ] `DELETE /artifacts/{id}` — Удаление артефакта
- [ ] `GET /artifacts/portfolio/conditions` — Условия загрузки портфолио
- [ ] `GET /artifacts/photo/conditions` — Условия загрузки фотографий

### Подсказки (suggests)

- [ ] `GET /suggests/skill_set` — Подсказки по ключевым навыкам
- [ ] `GET /suggests/areas` — Подсказки по регионам
- [ ] `GET /suggests/positions` — Подсказки по должностям резюме
- [ ] `GET /suggests/companies` — Подсказки по зарегистрированным организациям
- [ ] `GET /suggests/educational_institutions` — Подсказки по названиям учебных заведений
- [ ] `GET /suggests/professional_roles` — Подсказки по профессиональным ролям
- [ ] `GET /suggests/vacancy_search_keyword` — Подсказки по ключевым словам поиска вакансий
- [ ] `GET /suggests/resume_search_keyword` — Подсказки по ключевым словам поиска резюме
- [ ] `GET /suggests/area_leaves` — Подсказки по регионам (листьям дерева регионов)
- [ ] `GET /suggests/fields_of_study` — Подсказки по специализациям

### Справочники и dictionaries

- [ ] `GET /dictionaries` — Справочник полей и сущностей
- [ ] `GET /areas` — Справочник регионов
- [ ] `GET /professional_roles` — Справочник профессиональных ролей
- [ ] `GET /skills` — Справочник ключевых навыков
- [ ] `GET /metro` — Справочник станций метро
- [ ] `GET /educational_institutions` — Основная информация об учебных заведениях
- [ ] `GET /educational_institutions/{id}/faculties` — Список факультетов учебного заведения
- [ ] `GET /industries` — Справочник отраслей
- [ ] `GET /metro/{city_id}` — Список станций метро в указанном городе
- [ ] `GET /specializations` — Справочник специализаций
- [ ] `GET /employer_types` — Справочник типов работодателей
- [ ] `GET /areas/{area_id}` — Информация о регионе
- [ ] `GET /locales` — Справочник доступных локалей
- [ ] `GET /languages` — Справочник языков общения

### Статистика и аналитика

- [ ] `GET /vacancies/{vacancy_id}/stats` — Просмотр статистики по вакансии
- [ ] `GET /resumes/{resume_id}/views` — Просмотр статистики просмотров резюме
- [ ] `GET /salary_statistics/salary/evaluation` — Оценка заработной платы без прогноза
- [ ] `GET /clickme/statistics` — Получение информации о статистике рекламных кампаний в Clickme

### Информация о пользователе

- [ ] `GET /me` — Информация о текущем пользователе
- [ ] `POST /me` — Редактирование информации авторизованного пользователя
- [ ] `GET /applicant/comments` — Получение комментариев о соискателе
- [ ] `POST /applicant/comments` — Добавление комментария к соискателю

### Услуги и оплата

- [ ] `GET /payable/api_access` — Проверка доступа к платным методам API
- [ ] `GET /employers/{employer_id}/services/active` — Список активных услуг работодателя
- [ ] `GET /employers/{employer_id}/services/inactive` — Список неактивных услуг работодателя
- [ ] `POST /employers/{employer_id}/services/{service_id}` — Активация услуги

### Webhook API

- [ ] `POST /webhook/subscriptions` — Создать подписку на уведомление
- [ ] `GET /webhook/subscriptions` — Получить список уведомлений, на которые подписан пользователь
- [ ] `PUT /webhook/subscriptions/{subscription_id}` — Обновить подписку на уведомление
- [ ] `DELETE /webhook/subscriptions/{subscription_id}` — Удалить подписку на уведомление
