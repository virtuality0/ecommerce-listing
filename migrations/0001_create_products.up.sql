CREATE TABLE product(
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  price NUMERIC(
    12,
    2
  ) NOT NULL CHECK(price > 0),
  stock INTEGER NOT NULL DEFAULT 0
);
