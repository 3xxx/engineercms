var codeHandler = (function(){
    var base64Chars = [
            'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
            'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
            'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
            'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f',
            'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
            'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
            'w', 'x', 'y', 'z', '0', '1', '2', '3',
            '4', '5', '6', '7', '8', '9', '+', '/'
        ],
        encode = {
            'base64':codeBase64
        },
        decode = {
            'base64':decodeBase64
        }
    handleFormat = {
        'utf-8':toUTF8Binary
    };
    function stringToBinary(str , size , encodeType ){
        //  str-字符串 , size - 转换后的二进制位数 ,encodeType - 采用什么格式去保存二进制编码
        var i,
            len,
            binary = '';
        for ( i = 0 , len = str.length ; i < len ; i++ ){
            binary = binary + handleFormat[encodeType.toLowerCase()](str.charCodeAt(i));
        }
        //console.log(binary);
        return binary;
    }
    // 转换为以UTF-8格式的二进制数据
    function toUTF8Binary(unicode){
        var len,
            binary = '',
            star = 0,
            bitStream = unicode.toString(2), // 转换为二进制比特流
            bitLen = bitStream.length,
            i;
        if( unicode >= 0x000000 && unicode <= 0x00007F ){
            binary = bitStream;
            for( i = 0 , len = 8 ; i　< len-bitLen ; i ++ ){
                binary = 0 +binary; // 不足8位补0
            }
        }else if( unicode >=0x000080 && unicode <=0x0007FF ){
            binary = bitStream;
            for( i = 0 , len = 11 ; i　< len-bitLen ; i ++ ){
                binary = 0 +binary; // 不足11位补0
            }
            binary = '110'+binary.substr(0,5) + '10' + binary.substr(5,6);
        }
        else if( unicode >=0x000800 && unicode <=0x00FFFF ){
            binary = bitStream;
            for( i = 0 , len = 16 ; i　< len-bitLen ; i ++ ){
                binary = 0 +binary; // 不足16位补0
            };
            binary = '1110' + 
                     binary.substr(0,4) + 
                     '10' + 
                     binary.substr(4,6) + 
                     '10' + 
                     binary.substr(10,6);
        }
        else if( unicode >=0x010000 && unicode <=0x10FFFF ){
            binary = bitStream;
            for( i = 0 , len = 21 ; i　< len-bitLen ; i ++ ){
                binary = 0 +binary; // 不足21位补0
            }
            binary = '11110' + 
                     binary.substr(0,3) + 
                     '10' + 
                     binary.substr(3,6) + 
                     '10' + 
                     binary.substr(9,6) +
                     '10' + 
                     binary.substr(15,6);
        }
        return binary;
    }
    // 编码成base64格式
    function base64Parse(binary24,flag){
        var i,
            len,
            result = '',
            decode;
        if(flag == 1){
            for( i = 0 ; i < 4 ; i++){
                decode = parseInt(binary24.substr(i*6,6),2);
                result = result + base64Chars[decode];
            }
        }
        else{
            for ( i=0 , len = Math.floor(flag/6) ;i<len+1; i++){
                decode = parseInt(binary24.substr(i*6,6),2);
                result = result + base64Chars[decode];
            }
            for( i = 0; i < 3-len ;i ++){
                result = result + '=';
            }
        }
        return result;
    }
    // 解析为base64格式的二进制数据
    function codeBase64(str){
        var i,
            len,
            rem,
            mer,
            result = '',
            strBinaryAry = [],
            binary = stringToBinary(str , 8 , 'utf-8'); // base64是基于utf-8格式保存的二进制数据转换的
        len = binary.length;
        mer = Math.floor(len / 24);
        rem = len % 24;
        for( i = 0 ; i < mer ; i++){
            result = result +  base64Parse(binary.substr(i*24,24),1);
        }
        remCode = binary.substr(len-rem,rem);
        if( rem > 0 ){
            for( i =0 ; i < 24-rem ; i++){
                remCode = remCode + 0;
            }
            result = result +  base64Parse(remCode,rem)
        }
        return result;

    }
    // 解码base64格式的数据
    function decodeBase64(str){
        var i,
            j,
            k,
            len,
            t = 0,
            curbinary,
            start  = 0 ,
            flag = [
                {
                    str:'0',
                    len:8
                },
                {
                    str:'110',
                    len:11
                },
                {
                    str:'1110',
                    len:16
                },
                {
                    str:'11110',
                    len:21
                }],
            binary= '',
            newStr = '';
        for( i = 0 , len = str.length ; i < len ; i++){
            var curbinary  = base64Chars.indexOf(str.charAt(i)).toString(2);
            if( curbinary != '-1'){

                for( j = 0 ; curbinary.length <6 ; j++){
                    curbinary = 0 + curbinary;
                }
                binary = binary + curbinary;
            }
            if( i >= len-2 && str.charAt(i) == '='){
                ++t;
            }
        }
        if( t == 0 ){
            len = binary.length;
        }
        else{
            len = binary.length - (6-2*t)
        }

        for( ; start < len ;){
            for( j  = 0 ; j < 4 ; j++){

                if(binary.indexOf( flag[j].str ,start) == start){
                    if(flag[j].len == 8){
                        newStr = newStr +
                                 String.fromCharCode(parseInt(binary.substr(start,8),2));
                    }
                    else if(flag[j].len == 11){
                        newStr = newStr +
                                 String.fromCharCode(parsetInt(binary.substr(start+3,5) + 
                                 binary.substr(start+10,6),2));
                    }
                    else if(flag[j].len == 16){
                        newStr = newStr +     
                                 String.fromCharCode(parsetInt(binary.substr(start+4,4) + 
                                 binary.substr(start+10,6) + 
                                 binary.substr(start+18,6),2));
                    }
                    else if(flag[j].len == 21){
                        newStr = newStr + 
                                 String.fromCharCode(parseInt(binary.substr(start+5,3) + 
                                 binary.substr(start+10,6) + binary.substr(start+18,6) +
                                 binary.substr(start+26,6),2));
                    }
                    start  =  start  + flag[j].len;
                    break;
                }
            }
        }
        binary = null;
        return newStr;
    }
    return {
        encode:function(str ,type){
            return encode[type](str);
        },
        decode:function(str, type){
            return decode[type](str);
        }
    };
})();