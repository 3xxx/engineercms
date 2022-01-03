function Decode(e) {
    if (!e)
        return e;
    for (var t = "", o = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".indexOf(e[0]), i = 1; i < e.length - 2; i++) {
        for (var a = c(++o), s = e.charCodeAt(i), l = ""; /[0-9-]/.test(e[i + 1]); )
            l += e[++i];
        s = d(s, a, l = parseInt(l, 10) || 0);
        s ^= o - 1 & 31;
        t += String.fromCharCode(s);
    }
    return t
}

function Encode(text)
{
    var o = 53, s = "", result = "", prefix = 53; 
    for (var i = 0; i < text.length; i++)
    {
        var a = c(++o);
        s = text[i].charCodeAt();
        s ^= ((o - 1) & 31);
        s = d_reverse(s, a, 0);
        result += String.fromCharCode(s);
    } 
    return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".charAt(prefix) + result + "==";
}

function c(e) {
    for (var t = e.toString(), n = 0, r = 0; r < t.length; r++)
        n += parseInt(t.charAt(r), 10);
    return 10 < n ? n % 9 + 1 : n
}

function d(e, t, n) {
    for (var r = Math.abs(n); 0 < r--; )
        e -= t;
    return n < 0 && (e += 123), e
}

function d_reverse(e, t, n) {
    for (var r = Math.abs(n); 0 < r--; )
        e += t;
    return e
}

var license_data = "V3|BOARD4ALL|WILDCARD_ACTIVATION_KEY|2099";
var enc_license = Encode(license_data);

console.log("\n\x1b[32mLICENSE DATA:\x1b[0m \x1b[36m" + license_data + "\x1b[0m");
console.log("\n\x1b[32mLICENSE KEY:\x1b[0m \x1b[36m" + enc_license.replace(/\\/g, '\\\\') + "\x1b[0m");