// An Elegant Pairing Function by Matthew Szudzik @ Wolfram Research, Inc.
export function elegantPair(x, y) {
    return (x >= y) ? (x * x + x + y) : (y * y + x)
}