# Подготовка к Заданию 1 (SQL / Базы данных) — Полный справочник паттернов (Бакалавриат)

Задание 1 в направлении **«Программирование и информационные технологии» (ПиИТ)** для категории **«Бакалавриат»** может содержать самые разнообразные задачи по SQL — от оконных функций и рекурсии до сложных алгоритмов на графах и задач на пропуски и островки (Gaps & Islands).

В этом файле собран **полный энциклопедический справочник паттернов**, которые могут встретиться в Задании 1, с примерами кода и готовыми стратегиями решения.

---

## 📁 Файлы вариантов бакалавриата в репозитории

* 📄 [`ПиИТ_Демоверсия_Бакалавриат_Полуфинал_25-26.pdf`](file:///home/meidori/tbank-internship-preparation/yaprofi/it/25-26/ПиИТ_Демоверсия_Бакалавриат_Полуфинал_25-26.pdf) — *Паттерн 1: Оконные функции и EXCLUDE CURRENT ROW*
* 📄 [`ПиИТ_Демоверсия_Бакалавриат_23-24.pdf`](file:///home/meidori/tbank-internship-preparation/yaprofi/it/23-24/ПиИТ_Демоверсия_Бакалавриат_23-24.pdf) — *Паттерн 2: Рекурсивный граф и условие останова*
* 📄 [`ПиИТ_Демоверсия_Бакалавриат_22-23.pdf`](file:///home/meidori/tbank-internship-preparation/yaprofi/it/22-23/ПиИТ_Демоверсия_Бакалавриат_22-23.pdf) — *Паттерн 2: Иерархии оргструктуры и step/level*

---

## 🚀 ПОЛНЫЙ КАТАЛОГ ПАТТЕРНОВ ДЛЯ ЗАДАНИЯ 1

---

### 🧩 Паттерн 1. Оконные функции $\iff$ Подзапросы и JOINs

Оконные функции вычисляют результаты по окну строк. В олимпиаде часто требуют **переписать оконную функцию без `OVER`** или с лимитом на число `SELECT`.

#### 1.1. Агрегат без текущей строки (`EXCLUDE CURRENT ROW`)
* **Суть**: Найти сумму/среднее по группе, исключая текущий элемент.
* **Математика**: $\text{Result} = \text{TotalSum(Group)} - \text{CurrentValue}$
* **Код (без `OVER`, ровно 2 `SELECT`)**:
  ```sql
  SELECT "TrackId",
         "Milliseconds",
         (SELECT SUM("Milliseconds") FROM "Track" WHERE "GenreId" = $genre) - "Milliseconds" AS "Result"
  FROM "Track"
  WHERE "GenreId" = $genre
  ORDER BY "TrackId";
  ```

#### 1.2. Поиск Top-N элементов в группе без `ROW_NUMBER() / DENSE_RANK()`
* **Суть**: Получить первенцев/топ-3 сотрудников с максимальной зарплатой в каждом отделе.
* **Решение через коррелированный подзапрос**:
  ```sql
  SELECT e1."DepartmentId", e1."Name", e1."Salary"
  FROM "Employee" e1
  WHERE (
      SELECT COUNT(DISTINCT e2."Salary")
      FROM "Employee" e2
      WHERE e2."DepartmentId" = e1."DepartmentId" 
        AND e2."Salary" > e1."Salary"
  ) < 3
  ORDER BY e1."DepartmentId", e1."Salary" DESC;
  ```

#### 1.3. Бегущий итог (Running Total) / Скользящие средние (Moving Average)
* **С оконной функцией**: `SUM(val) OVER (PARTITION BY grp ORDER BY dt ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)`
* **Без оконной функции (Self-JOIN)**:
  ```sql
  SELECT t1."Id", t1."Dt", SUM(t2."Val") AS "RunningTotal"
  FROM "Transactions" t1
  JOIN "Transactions" t2 ON t1."Grp" = t2."Grp" AND t2."Dt" <= t1."Dt"
  GROUP BY t1."Id", t1."Dt"
  ORDER BY t1."Dt";
  ```

#### 1.4. `LAG()` / `LEAD()` без специальных функций
* **Суть**: Получить значение из предыдущей или следующей строки.
* **Решение через Self-JOIN по предшествующей дате/индексу**:
  ```sql
  SELECT t1."Id", t1."Dt", t1."Val", MAX(t2."Val") AS "PrevVal"
  FROM "Events" t1
  LEFT JOIN "Events" t2 
         ON t1."GroupId" = t2."GroupId" 
        AND t2."Dt" < t1."Dt"
  GROUP BY t1."Id", t1."Dt", t1."Val";
  ```

---

### 🌳 Паттерн 2. Рекурсия, Графы и Деревья (`WITH RECURSIVE`)

#### 2.1. Иерархия (Предки и Потомки)
* **Суть**: Получить всех начальников/подчиненных или всё дерево категорий.
* **Шаблон**:
  ```sql
  WITH RECURSIVE Tree AS (
      -- 1. Якорь: стартовый узел
      SELECT "Id", "ParentId", 1 AS "Level"
      FROM "Categories"
      WHERE "Id" = $start_id
      
      UNION ALL
      
      -- 2. Рекурсивный шаг: соединение с потомками
      SELECT c."Id", c."ParentId", t."Level" + 1
      FROM "Categories" c
      JOIN Tree t ON c."ParentId" = t."Id"
  )
  SELECT * FROM Tree;
  ```

#### 2.2. Поиск путей в графе и обнаружение циклов (Cycle Detection)
* **Суть**: Найти путь между узлами `A` и `B` в графе и избежать бесконечного цикла.
* **Шаблон с аккумулятором пути (String Path)**:
  ```sql
  WITH RECURSIVE Paths AS (
      SELECT "FromNode", "ToNode", CAST("FromNode" AS TEXT) || '->' || CAST("ToNode" AS TEXT) AS "Path"
      FROM "Edges"
      WHERE "FromNode" = $start_node
      
      UNION ALL
      
      SELECT e."FromNode", e."ToNode", p."Path" || '->' || CAST(e."ToNode" AS TEXT)
      FROM "Edges" e
      JOIN Paths p ON e."FromNode" = p."ToNode"
      WHERE p."Path" NOT LIKE '%' || CAST(e."ToNode" AS TEXT) || '%' -- Защита от циклов
  )
  SELECT * FROM Paths WHERE "ToNode" = $target_node;
  ```

#### 2.3. Генерация последовательностей и рядов дат (Sequence Generator)
* **Суть**: Сгенерировать список чисел от 1 до N или непрерывный календарь дат.
  ```sql
  WITH RECURSIVE Calendar AS (
      SELECT DATE('2026-01-01') AS "Date"
      UNION ALL
      SELECT DATE("Date", '+1 day')
      FROM Calendar
      WHERE "Date" < '2026-01-31'
  )
  SELECT * FROM Calendar;
  ```

---

### 🏝 Паттерн 3. Пропуски и Островки (Gaps & Islands)

Один из самых частых продвинутых алгоритмических паттернов в SQL.

#### 3.1. Нахождение непрерывных диапазонов (Islands)
* **Задача**: Найти периоды непрерывных дней, когда пользователь заходил на сайт.
* **Магический трюк**: Если вычесть `ROW_NUMBER()` из даты, то у всех дней одного непрерывного островка полученная разность будет **одинаковой**!
  ```sql
  WITH Ranked AS (
      SELECT "UserId", "LogDate",
             DATE("LogDate", '-' || ROW_NUMBER() OVER (PARTITION BY "UserId" ORDER BY "LogDate") || ' day') AS "Grp"
      FROM "UserActivity"
  )
  SELECT "UserId", MIN("LogDate") AS "StartDate", MAX("LogDate") AS "EndDate", COUNT(*) AS "DaysCount"
  FROM Ranked
  GROUP BY "UserId", "Grp"
  HAVING COUNT(*) >= 3;
  ```

#### 3.2. Поиск пропущенных значений/дат (Gaps)
* **Задача**: Найти пропущенные ID заказов или пропущенные дни в логах.
* **Решение 1 (через `LEFT JOIN` с генератором)**:
  ```sql
  SELECT c."Date" AS "MissingDate"
  FROM Calendar c
  LEFT JOIN "Logs" l ON c."Date" = l."LogDate"
  WHERE l."LogDate" IS NULL;
  ```
* **Решение 2 (через сравнение со следующей строкой)**:
  ```sql
  SELECT "Id" + 1 AS "GapStart", "NextId" - 1 AS "GapEnd"
  FROM (
      SELECT "Id", LEAD("Id") OVER (ORDER BY "Id") AS "NextId"
      FROM "Orders"
  ) t
  WHERE "NextId" - "Id" > 1;
  ```

---

### ➗ Паттерн 4. Реляционная алгебра и Реляционное деление (Relational Division)

#### 4.1. Реляционное деление («Все элементы из набора»)
* **Задача**: Найти покупателей, которые купили **абсолютно все** товары из категории 'Электроника'.
* **Способ 1 (через `GROUP BY + HAVING COUNT(DISTINCT)`)**:
  ```sql
  SELECT "CustomerId"
  FROM "Purchases" p
  JOIN "Products" prod ON p."ProductId" = prod."Id"
  WHERE prod."Category" = 'Electronics'
  GROUP BY "CustomerId"
  HAVING COUNT(DISTINCT p."ProductId") = (
      SELECT COUNT(*) FROM "Products" WHERE "Category" = 'Electronics'
  );
  ```
* **Способ 2 (через двойное `NOT EXISTS` — классическая реляционная алгебра)**:
  «Найти покупателя, для которого НЕ существует такого товара в 'Electronics', который он НЕ купил».
  ```sql
  SELECT c."CustomerId"
  FROM "Customers" c
  WHERE NOT EXISTS (
      SELECT p."Id" FROM "Products" p WHERE p."Category" = 'Electronics'
      EXCEPT
      SELECT pur."ProductId" FROM "Purchases" pur WHERE pur."CustomerId" = c."CustomerId"
  );
  ```

#### 4.2. Точное совпадение множеств (Set Equality)
* **Задача**: Найти пользователей, имеющих ровно такой же набор ролей, как у 'User_42'.
  ```sql
  SELECT "UserId"
  FROM "UserRoles"
  WHERE "UserId" <> 42
  GROUP BY "UserId"
  HAVING COUNT(*) = (SELECT COUNT(*) FROM "UserRoles" WHERE "UserId" = 42)
     AND COUNT(CASE WHEN "RoleId" IN (SELECT "RoleId" FROM "UserRoles" WHERE "UserId" = 42) THEN 1 END) = (SELECT COUNT(*) FROM "UserRoles" WHERE "UserId" = 42);
  ```

---

### 🔄 Паттерн 5. Поворот таблиц (Pivot) и Транспонирование

#### 5.1. Pivot (Строки в Столбцы) без специализированных операторов
* **Задача**: Преобразовать таблицу продаж по месяцам `(Year, Month, Sales)` в вид `(Year, Jan_Sales, Feb_Sales, Mar_Sales)`.
* **Решение через условную агрегацию (Conditional Aggregation)**:
  ```sql
  SELECT "Year",
         SUM(CASE WHEN "Month" = 1 THEN "Sales" ELSE 0 END) AS "Jan_Sales",
         SUM(CASE WHEN "Month" = 2 THEN "Sales" ELSE 0 END) AS "Feb_Sales",
         SUM(CASE WHEN "Month" = 3 THEN "Sales" ELSE 0 END) AS "Mar_Sales"
  FROM "MonthlySales"
  GROUP BY "Year"
  ORDER BY "Year";
  ```

#### 5.2. Unpivot (Столбцы в Строки)
* **Задача**: Преобразовать столбцы `(Id, Q1, Q2, Q3)` обратно в строки `(Id, Quarter, Value)`.
* **Решение через `UNION ALL`**:
  ```sql
  SELECT "Id", 'Q1' AS "Quarter", "Q1" AS "Value" FROM "QuarterlyData"
  UNION ALL
  SELECT "Id", 'Q2' AS "Quarter", "Q2" AS "Value" FROM "QuarterlyData"
  UNION ALL
  SELECT "Id", 'Q3' AS "Quarter", "Q3" AS "Value" FROM "QuarterlyData";
  ```

---

### ⚡ Паттерн 6. Сложные соединение по диапазонам (Non-Equi JOINs)

#### 6.1. Соединение по интервалам/диапазонам
* **Задача**: Сопоставить транзакции пользователя с действующим на момент транзакции скидочным купоном `(CouponStart, CouponEnd)`.
  ```sql
  SELECT t."TransactionId", t."Amount", c."DiscountRate"
  FROM "Transactions" t
  LEFT JOIN "Coupons" c 
         ON t."UserId" = c."UserId" 
        AND t."TxDate" BETWEEN c."StartDate" AND c."EndDate";
  ```

#### 6.2. Эмуляция `FULL OUTER JOIN` в СУБД без его поддержки (например, в SQLite)
  ```sql
  SELECT t1."Id", t1."Val1", t2."Val2"
  FROM "TableA" t1
  LEFT JOIN "TableB" t2 ON t1."Id" = t2."Id"
  
  UNION ALL
  
  SELECT t2."Id", NULL AS "Val1", t2."Val2"
  FROM "TableB" t2
  WHERE NOT EXISTS (SELECT 1 FROM "TableA" t1 WHERE t1."Id" = t2."Id");
  ```

---

## 📋 Итоговый чек-лист подготовки ко ВСЕМ форматам Задания 1

- [ ] **Оконные функции**: `OVER`, `PARTITION BY`, `ROWS/RANGE BETWEEN`, `EXCLUDE CURRENT ROW/GROUP/TIES`.
- [ ] **Рекурсия CTE**: `WITH RECURSIVE`, якорь, шаг `UNION ALL`, глубина `step`, поиск циклов `NOT LIKE`.
- [ ] **Gaps & Islands**: Разность `date - ROW_NUMBER()`, поиск пропущенных ID через `LEAD()`.
- [ ] **Реляционное деление**: Поиск «всех предметов» через `HAVING COUNT(DISTINCT)` и `NOT EXISTS`.
- [ ] **Pivot/Unpivot**: Условный агрегат `SUM(CASE WHEN ...)`, разворачивание через `UNION ALL`.
- [ ] **Эквивалентные переписывания**:
  - `OVER ()` $\rightarrow$ подзапросы/Self-JOIN.
  - `LAG/LEAD` $\rightarrow$ `LEFT JOIN` по предшествующему значению.
  - Ограничение на **число операторов `SELECT`** (ровно 2 SELECT).
- [ ] **Синтаксис и оформление**: Идентификаторы в двойных кавычках `"Column"`, параметры с верным синтаксисом `$var`, правильная сортировка `ORDER BY`.
