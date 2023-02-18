import _sodium from 'libsodium-wrappers';

const AUTH_CONTEXT = 'AUTH_KEY';

// Register
export const generateRequiredKeysAsync = async (password: string) => {
    await _sodium.ready;
    const sodium = _sodium;

    const salt = sodium.randombytes_buf(sodium.crypto_pwhash_SALTBYTES);
    const passwordHash = sodium.crypto_pwhash(
        sodium.crypto_box_SEEDBYTES,
        password,
        salt,
        sodium.crypto_pwhash_OPSLIMIT_INTERACTIVE,
        sodium.crypto_pwhash_MEMLIMIT_INTERACTIVE,
        sodium.crypto_pwhash_ALG_DEFAULT
    );

    const authKey = sodium.crypto_kdf_derive_from_key(
        32,
        1,
        AUTH_CONTEXT,
        passwordHash
    );

    const masterKey = sodium.crypto_kdf_keygen();

    const nonce = sodium.randombytes_buf(sodium.crypto_secretbox_NONCEBYTES);
    const encMasterKey = sodium.crypto_secretbox_easy(
        masterKey,
        nonce,
        passwordHash
    );

    return {
        salt: sodium.to_base64(salt),
        nonce: sodium.to_base64(nonce),
        encMasterKey: sodium.to_base64(encMasterKey),
        authKey: sodium.to_base64(authKey),
    };
};

// After submitting email when logging in, you'll get back a hash, which can be passed to this function
export const retrieveAuthKeyAndPasswordHashAsync = async (
    password: string,
    salt: string
) => {
    await _sodium.ready;
    const sodium = _sodium;

    const uIntSalt = sodium.from_base64(salt);
    const passwordHash = sodium.crypto_pwhash(
        sodium.crypto_box_SEEDBYTES,
        password,
        uIntSalt,
        sodium.crypto_pwhash_OPSLIMIT_INTERACTIVE,
        sodium.crypto_pwhash_MEMLIMIT_INTERACTIVE,
        sodium.crypto_pwhash_ALG_DEFAULT
    );

    const authKey = sodium.crypto_kdf_derive_from_key(
        32,
        1,
        AUTH_CONTEXT,
        passwordHash
    );

    return {
        passwordHash,
        authKey: sodium.to_base64(authKey),
    };
};

// After submitting the auth key when logging in, you'll get back a packet, if login was successful, which can be passed to this function
export const decryptMasterKeyAsync = async (
    passwordHash: Uint8Array,
    nonce: string,
    encMasterKey: string
) => {
    await _sodium.ready;
    const sodium = _sodium;

    const uIntNonce = sodium.from_base64(nonce); // 24 Bytes
    const decMasterKey = sodium.crypto_secretbox_open_easy(
        encMasterKey,
        uIntNonce,
        passwordHash
    );

    return decMasterKey;
};
