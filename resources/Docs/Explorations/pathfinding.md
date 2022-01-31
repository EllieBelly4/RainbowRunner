Terrain definitions contain CollisionObject = X which point to models that contain walkable and non-walkable meshes.

Townston:

Following contain walkable surfaces for townston

Townston_tier_1
Townston_tier_2
Townston_tier_3

## WorldToGridPos

```
bool __userpurge PathMap::WorldToGridPos@<al>(#310 *this@<ecx>, int a2@<eax>, int *a3@<edi>, _DWORD *a4@<esi>, int a5, int a6, int *a7, int *a8)
{
  int v8; // eax

  *a3 = (a2 - 10 * a4[6]) / 10;
  v8 = ((int)this - 10 * a4[9]) / 10;
  *(_DWORD *)a5 = v8;
  return *a3 >= 0 && *a3 < a4[10] && v8 >= 0 && v8 < a4[11];
```

## Things

Town does not have "PlaceholderStatics"