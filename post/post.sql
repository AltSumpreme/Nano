CREATE TABLE IF NOT EXISTS posts{
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    account_id UUID REFERENCES accounts(id),

}

CREATE TABLE comments{
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(id),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    account_id UUID REFERENCES accounts(id),
}

CREATE TABLE likes{
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts(id),
    account_id UUID REFERENCES accounts(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (post_id, account_id)
}