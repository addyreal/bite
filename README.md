# bite
WIP, TODO:
allow adding patterns beginning with wildcard

Provides a simple interface for implementing byte signature prefix lookup supporting wildcards.

## Exports

- Init
- Add
- Get
- Find

<hr>

### `Init() *temp`
**Description:** Initializes a handle for the builder.

**Returns:**
- `*temp`

<hr>

### `(*temp) Add(res any, pattern ...int)`
**Description:** Adds a pattern supporting wildcards, which, when looked up, returns `res`.

**Args:**
- `res`
- `pattern` - a slice of `[0,255]` for exact match, `-1` for wildcard

<hr>

### `(*temp) Get() *table`
**Description:** Builds and returns the lookup table.

**Returns:**
- `*table`

<hr>

### `(*table) Find(b []byte) (any, bool)`
**Description:** Tries to match the longest pattern provided matching the beginning of `b`, after which it returns the desired result and lookup success.

**Args:**
- `b`

**Returns:**
- `any` - result
- `bool` - lookup success
