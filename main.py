import time
from pythonicsql import dialects, go


class PythonicSQL:
    def __init__(self, dialect: str, uri: str):
        self.client = dialects.new_client(dialect=dialect, uri=uri)


s = time.time()
p = PythonicSQL(
    dialect="postgres",
    uri="postgres://dev-duckorm:postgres123@localhost:5432/dev-duckorm?sslmode=disable",
)

p.client.connect()

res = p.client.builder.select(go.Slice_string(["id"])).from_("interaction").exec()

for r in res:
    print(r)

e = time.time()
print((e - s) * 1000)
