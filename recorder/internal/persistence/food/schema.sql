CREATE TABLE food (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT null,
    description TEXT NOT null
);

CREATE TABLE foodRecord (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    userId UUID NOT null,
    weight float(5)
);