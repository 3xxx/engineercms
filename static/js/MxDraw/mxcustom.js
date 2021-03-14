
// �����Զ���ʵ�庯��
function InsertCustomEntity() {

    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "��ȡ��һ��";
    if (getPt.go() != 1)
        return;
    var frstPt = getPt.value();
    if (frstPt == null)
        return;

    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");

    getSecondPt.message = "��ȡ�ڶ���";
    getSecondPt.basePoint = frstPt;

    getSecondPt.setUseBasePt(true);


    if (getSecondPt.go() != 1)
        return;

    var secondPt = getSecondPt.value();
    if (secondPt == null)
        return;

    var ent = mxOcx.DrawCustomEntity("TestMxCustomEntity", "");


    ent.SetPoint("spt", frstPt);
    ent.SetPoint("ept", secondPt);
	ent.SetString("MxObjectAppName", "�ҵ��Զ���ʵ����");
}

// ��ʽ���ַ���
function formatNumber(num, pattern) {
    var strarr = num ? num.toString().split('.') : ['0'];
    var fmtarr = pattern ? pattern.split('.') : [''];
    var retstr = '';

    // ��������
    var str = strarr[0];
    var fmt = fmtarr[0];
    var i = str.length - 1;
    var comma = false;
    for (var f = fmt.length - 1; f >= 0; f--) {
        switch (fmt.substr(f, 1)) {
            case '#':
                if (i >= 0) retstr = str.substr(i--, 1) + retstr;
                break;
            case '0':
                if (i >= 0) retstr = str.substr(i--, 1) + retstr;
                else retstr = '0' + retstr;
                break;
            case ',':
                comma = true;
                retstr = ',' + retstr;
                break;
        }
    }
    if (i >= 0) {
        if (comma) {
            var l = str.length;
            for (; i >= 0; i--) {
                retstr = str.substr(i, 1) + retstr;
                if (i > 0 && ((l - i) % 3) == 0) retstr = ',' + retstr;
            }
        }
        else retstr = str.substr(0, i + 1) + retstr;
    }

    retstr = retstr + '.';
    // ����С������
    str = strarr.length > 1 ? strarr[1] : '';
    fmt = fmtarr.length > 1 ? fmtarr[1] : '';
    i = 0;
    for (var f = 0; f < fmt.length; f++) {
        switch (fmt.substr(f, 1)) {
            case '#':
                if (i < str.length) retstr += str.substr(i++, 1);
                break;
            case '0':
                if (i < str.length) retstr += str.substr(i++, 1);
                else retstr += '0';
                break;
        }
    }
    return retstr.replace(/^,+/, '').replace(/\.$/, '');
}


// �Զ���ʵ����ƺ���
/*
function ExplodeFun(pCustomEntity, pWorldDraw) {


    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;

        var mxUtility = mxOcx.NewUtility();
        var vec = ept.SumVector(stp);

        vec.Mult(0.5);

        var midPt = mxOcx.NewPoint();

        midPt.x = stp.x;
        midPt.y = stp.y;
        midPt.Add(vec);

        var dAng = vec.Angle();
        dAng = mxUtility.GetDimAngle(dAng);

        var dDis = 0.0;
        dDis = stp.DistanceTo(ept);

        var sTxt = "L=" + formatNumber(dDis, '#.##');

        dAng = dAng * 180.0 / 3.14159265;

        vec.RotateByXyPlan(3.14159265 / 2.0);
        vec.Normalize();
        vec.Mult(10);

        stp.Add(vec);
        ept.Add(vec);

        pWorldDraw.DrawLine(stp.x, stp.y, ept.x, ept.y);

        vec.Mult(2);

        stp.Sum(vec);
        ept.Sum(vec);

        pWorldDraw.DrawLine(stp.x, stp.y, ept.x, ept.y);

        pWorldDraw.SetColorIndex(1);

        pWorldDraw.DrawText(midPt.x, midPt.y, sTxt, 5, dAng,
			1, 2);

        mxOcx.SetEventRet(1);

    }
}
*/

// �����Զ���ʵ��е�
/*function GetGripPointsFun(pCustomEntity) {

    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;

        var ret = mxOcx.NewResbuf();

        ret.AddPoint(stp);
        ret.AddPoint(ept);

        mxOcx.SetEventRetEx(ret);
    }
}*/
// �ƶ��Զ���ʵ��е�
/*function MoveGripPointsFun(pCustomEntity, lGridIndex, dOffsetX, dOffsetY) {

    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;


        if (lGridIndex == 0) {
            stp.x = stp.x + dOffsetX;
            stp.y = stp.y + dOffsetY;
            pCustomEntity.SetPoint("spt", stp);
        }
        else {
            ept.x = ept.x + dOffsetX;
            ept.y = ept.y + dOffsetY;
            pCustomEntity.SetPoint("ept", ept);
        }

        mxOcx.SetEventRet(1);
    }
}*/
// �任�Զ���ʵ��
function TransformByFun(pCustomEntity, pMatXform) {

    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;


        stp.TransformBy(pMatXform);
        ept.TransformBy(pMatXform);

        pCustomEntity.SetPoint("spt", stp);

        pCustomEntity.SetPoint("ept", ept);
        mxOcx.SetEventRet(1);
    }
}
// �����Զ���ʵ����С���
function GetGeomExtentsFun(pCustomEntity) {

    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity") {
        if (!pCustomEntity.IsHave("ept"))
            return;

        var stp = pCustomEntity.GetPoint("spt");
        if (stp == null)
            return;

        var ept = pCustomEntity.GetPoint("ept");
        if (ept == null)
            return;

        var ret = mxOcx.NewResbuf();

        ret.AddPoint(stp);
        ret.AddPoint(ept);

        mxOcx.SetEventRetEx(ret);
    }
}

function GetOsnapPointsFun(pCustomEntity, lOsnapMode, dPickPointX, dPickPointY, dLastPointX, dLastPointY) {

    // 		enum OsnapMode         { kOsModeEnd          = 1,
    // 			kOsModeMid          = 2,
    // 			kOsModeCen          = 3,
    // 			kOsModeNode         = 4,
    // 			kOsModeQuad         = 5,
    // 			kOsModeIns          = 7,
    // 			kOsModePerp         = 8,
    // 			kOsModeTan          = 9,
    // 			kOsModeNear         = 10,
    // 			kOsModeInt          = 11
    // 		};

    var sGuid = pCustomEntity.Guid;
    if (sGuid == "TestMxCustomEntity")
    {

		if(!pCustomEntity.IsHave("ept") )
			return;

		var stp =  pCustomEntity.GetPoint("spt");
		if(stp == null)
			return;

		var ept =  pCustomEntity.GetPoint("ept");
		if(ept == null)
			return;


		if(lOsnapMode == 1)
		{
			// �˵㲶׽��
			var pickPoint = mxOcx.NewPoint();

			pickPoint.x = dPickPointX;
			pickPoint.y = dPickPointY;

			var dDis1 = pickPoint.DistanceTo(stp);
			var dDis2 = pickPoint.DistanceTo(ept);

             var ret = mxOcx.NewResbuf();

			if(dDis1 < dDis2)
			{
			    ret.AddDouble(stp.x);
			    ret.AddDouble(stp.y);
			}
			else {
			    ret.AddDouble(ept.x);
			    ret.AddDouble(ept.y);

			}

			mxOcx.SetEventRetEx(ret);
		}
		else if(lOsnapMode == 10)
		{
			// ����㲶׽.
			var pickPoint = mxOcx.NewPoint();

			pickPoint.x = dPickPointX;
			pickPoint.y = dPickPointY;


			var  line = mxOcx.NewEntity("IMxDrawLine");
			line.EndPoint   = ept;
			line.StartPoint = stp;

			var closePoint = line.GetClosestPointTo2(pickPoint,false);

			if(closePoint != null)
			{
			    var ret = mxOcx.NewResbuf();
			    ret.AddDouble(closePoint.x);
			    ret.AddDouble(closePoint.y);
			    mxOcx.SetEventRetEx(ret);
			}
		}

	}
}