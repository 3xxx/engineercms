    function popFileSelector() {
      $(`<input type="file" value="选择文件"></input>`).click().on("change", event => {
        let file = event.target.files[0];
        let file_reader = new FileReader();
        file_reader.onload = () => {
          let fc = file_reader.result;
          // console.log(fc); // 打印文件文本内容
          // var obj = fc.parseJSON(); //由JSON字符串转换为JSON对象
          var obj = JSON.parse(fc); //由JSON字符串转换为JSON对象
          loadData(obj);
        };
        file_reader.readAsText(file, "UTF-8");
      });
    }



    const readline = require("readline");

    const fs = require("fs");

    const Path = require("path");

    const DIR = "C:/Users/Admin/Desktop/";

    const dir = fs.readdirSync(DIR);

    dir.forEach(k => {
      if (k.indexOf(".kml") !== -1) {
        const path = DIR + k;
        const inp = readline.createInterface({
          input: fs.createReadStream(path)
        });
        let results = "";
        inp.on("line", line => {
          if (!line.length || line.startsWith(";")) return;
          if (line.indexOf("<Placemark") !== -1) {
            results += "{";
          }
          results += line;
          if (line.indexOf("</Placemark") !== -1) {
            results += "}";
          }
        });
        inp.on("close", () => {
          const bb = results.split("}");
          const geojson = {
            type: "FeatureCollection",
            features: []
          };
          const geojson2 = {
            type: "FeatureCollection",
            features: []
          };
          for (let i = 0; i < bb.length - 2; i++) {
            if (bb[i].split("<Placemark id=")[1]) {
              const coor = bb[i].split("<coordinates>")[1].split("</coordinates>")[0].split(",");
              const coor1 = bb[i + 1].split("<coordinates>")[1].split("</coordinates>")[0].split(",");
              geojson.features.push({
                type: "Feature",
                geometry: {
                  type: "LineString",
                  coordinates: [
                    [Number(coor[0]), Number(coor[1])],
                    [Number(coor1[0]), Number(coor1[1])]
                  ]
                },
                properties: {
                  measure: bb[i].split('<Placemark id="')[1].split('">')[0]
                }
              });
            }
          }
          for (let i = 0; i < bb.length; i++) {
            if (bb[i].split("<Placemark id=")[1]) {
              const coor = bb[i].split("<coordinates>")[1].split("</coordinates>")[0].split(",");
              geojson2.features.push({
                type: "Feature",
                geometry: {
                  type: "Point",
                  coordinates: [Number(coor[0]), Number(coor[1])]
                },
                properties: {
                  measure: bb[i].split('<Placemark id="')[1].split('">')[0]
                }
              });
            }
          }
          fs.writeFileSync(DIR + k.split(".kml")[0] + ".geojson", JSON.stringify(geojson, null, 2));
          fs.writeFileSync(DIR + k.split(".kml")[0] + "点.geojson", JSON.stringify(geojson2, null, 2));
        });
      }
    });
    // node generator.js