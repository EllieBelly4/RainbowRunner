import blowfish

AUTH_BLOWFISHKEY = b"[;'.]94-31==-%&@!^+]\x00"

bf = blowfish.Cipher(AUTH_BLOWFISHKEY, byte_order="little")

data_decrypted = b"".join(bf.decrypt_ecb(bytearray.fromhex("624A79FE01ED0600")))
print(data_decrypted)