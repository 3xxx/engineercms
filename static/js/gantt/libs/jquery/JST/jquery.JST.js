$.fn.loadTemplates = function() {
  $.JST.loadTemplates($(this));
  return this;
};

$.JST = {
  _templates: new Object(),
  _decorators:new Object(),

  loadTemplates: function(elems) {
    // var elems[0].innerHTML="↵<div class="__template__" type="GANTBUTTONS">↵  <!-- <div class="ganttButtonBar noprint">↵    <h1 class="ganttTitle" title="gantt"><img src="res/twGanttLogo.png" alt="gantt" align="absmiddle"><img src="res/twproject-badge.png" style="max-width: 120px" />↵</h1>↵    <div class="buttons">↵↵↵      <button onclick="$('#workSpace').trigger('undo.gantt');return false;" class="button textual icon requireCanWrite" title="undo"><span class="teamworkIcon">&#39;</span></button>↵      <button onclick="$('#workSpace').trigger('redo.gantt');return false;" class="button textual icon requireCanWrite" title="redo"><span class="teamworkIcon">&middot;</span></button>↵      <span class="ganttButtonSeparator requireCanWrite requireCanAdd"></span>↵      <button onclick="$('#workSpace').trigger('addAboveCurrentTask.gantt');return false;" class="button textual icon requireCanWrite requireCanAdd" title="insert above"><span class="teamworkIcon">l</span></button>↵      <button onclick="$('#workSpace').trigger('addBelowCurrentTask.gantt');return false;" class="button textual icon requireCanWrite requireCanAdd" title="insert below"><span class="teamworkIcon">X</span></button>↵      <span class="ganttButtonSeparator requireCanWrite requireCanInOutdent"></span>↵      <button onclick="$('#workSpace').trigger('outdentCurrentTask.gantt');return false;" class="button textual icon requireCanWrite requireCanInOutdent" title="un-indent task"><span class="teamworkIcon">.</span></button>↵      <button onclick="$('#workSpace').trigger('indentCurrentTask.gantt');return false;" class="button textual icon requireCanWrite requireCanInOutdent" title="indent task"><span class="teamworkIcon">:</span></button>↵      <span class="ganttButtonSeparator requireCanWrite requireCanMoveUpDown"></span>↵      <button onclick="$('#workSpace').trigger('moveUpCurrentTask.gantt');return false;" class="button textual icon requireCanWrite requireCanMoveUpDown" title="move up"><span class="teamworkIcon">k</span></button>↵      <button onclick="$('#workSpace').trigger('moveDownCurrentTask.gantt');return false;" class="button textual icon requireCanWrite requireCanMoveUpDown" title="move down"><span class="teamworkIcon">j</span></button>↵↵      <span class="ganttButtonSeparator requireCanWrite"></span>↵      <button onclick="$('#workSpace').trigger('deleteCurrentTask.gantt');return false;" class="button textual icon delete requireCanWrite" title="Delete"><span class="teamworkIcon">&cent;</span></button>↵↵      <span class="ganttButtonSeparator requireCanAddIssue"></span>↵      <button onclick="$('#workSpace').trigger('addIssue.gantt');return false;" class="button textual icon requireCanAddIssue" title="add issue / todo"><span class="teamworkIcon">i</span></button>↵↵↵      <span class="ganttButtonSeparator"></span>↵      <button onclick="$('#workSpace').trigger('expandAll.gantt');return false;" class="button textual icon " title="EXPAND_ALL"><span class="teamworkIcon">6</span></button>↵      <button onclick="$('#workSpace').trigger('collapseAll.gantt'); return false;" class="button textual icon " title="COLLAPSE_ALL"><span class="teamworkIcon">5</span></button>↵↵    <span class="ganttButtonSeparator"></span>↵      <button onclick="$('#workSpace').trigger('zoomMinus.gantt'); return false;" class="button textual icon " title="zoom out"><span class="teamworkIcon">)</span></button>↵      <button onclick="$('#workSpace').trigger('zoomPlus.gantt');return false;" class="button textual icon " title="zoom in"><span class="teamworkIcon">(</span></button>↵    <span class="ganttButtonSeparator"></span>↵      <button onclick="print();return false;" class="button textual icon " title="Print"><span class="teamworkIcon">p</span></button>↵    <span class="ganttButtonSeparator"></span>↵      <button onclick="ge.gantt.showCriticalPath=!ge.gantt.showCriticalPath; ge.redraw();return false;" class="button textual icon requireCanSeeCriticalPath" title="CRITICAL_PATH"><span class="teamworkIcon">&pound;</span></button>↵    <span class="ganttButtonSeparator requireCanSeeCriticalPath"></span>↵      <button onclick="ge.splitter.resize(.1);return false;" class="button textual icon" ><span class="teamworkIcon">F</span></button>↵      <button onclick="ge.splitter.resize(50);return false;" class="button textual icon" ><span class="teamworkIcon">O</span></button>↵      <button onclick="ge.splitter.resize(100);return false;" class="button textual icon"><span class="teamworkIcon">R</span></button>↵↵    <button onclick="editResources();" class="button textual requireWrite" title="edit resources"><span class="teamworkIcon">M</span></button>↵      &nbsp; &nbsp; &nbsp; &nbsp;↵    <button onclick="saveGanttOnServer();" class="button first big requireWrite" title="Save">Save</button>↵    <button onclick='newProject();' class='button requireWrite newproject'><em>clear project</em></button>↵    <button class="button login" title="login/enroll" onclick="loginEnroll($(this));" style="display:none;">login/enroll</butto…rt</label>&nbsp;&nbsp;&nbsp;&nbsp;↵            <input type="checkbox" id="startIsMilestone" name="startIsMilestone" value="yes"> &nbsp;<label for="startIsMilestone">is milestone</label>&nbsp;↵            <br><input type="text" name="start" id="start" size="8" class="formElements dateField validated date" autocomplete="off" maxlength="255" value="" oldvalue="1" entrytype="DATE">↵            <span title="calendar" id="starts_inputDate" class="teamworkIcon openCalendar" onclick="$(this).dateField({inputField:$(this).prevAll(':input:first'),isSearchField:false});">m</span>          </div>↵        </td>↵        <td nowrap="">↵          <label for="end">End</label>&nbsp;&nbsp;&nbsp;&nbsp;↵          <input type="checkbox" id="endIsMilestone" name="endIsMilestone" value="yes"> &nbsp;<label for="endIsMilestone">is milestone</label>&nbsp;↵          <br><input type="text" name="end" id="end" size="8" class="formElements dateField validated date" autocomplete="off" maxlength="255" value="" oldvalue="1" entrytype="DATE">↵          <span title="calendar" id="ends_inputDate" class="teamworkIcon openCalendar" onclick="$(this).dateField({inputField:$(this).prevAll(':input:first'),isSearchField:false});">m</span>↵        </td>↵        <td nowrap="" >↵          <label for="duration" class=" ">Days</label><br>↵          <input type="text" name="duration" id="duration" size="4" class="formElements validated durationdays" title="Duration is in working days." autocomplete="off" maxlength="255" value="" oldvalue="1" entrytype="DURATIONDAYS">&nbsp;↵        </td>↵      </tr>↵↵      <tr>↵        <td  colspan="2">↵          <label for="status" class=" ">status</label><br>↵          <select id="status" name="status" class="taskStatus" status="(#=obj.status#)"  onchange="$(this).attr('STATUS',$(this).val());">↵            <option value="STATUS_ACTIVE" class="taskStatus" status="STATUS_ACTIVE" >active</option>↵            <option value="STATUS_SUSPENDED" class="taskStatus" status="STATUS_SUSPENDED" >suspended</option>↵            <option value="STATUS_DONE" class="taskStatus" status="STATUS_DONE" >completed</option>↵            <option value="STATUS_FAILED" class="taskStatus" status="STATUS_FAILED" >failed</option>↵            <option value="STATUS_UNDEFINED" class="taskStatus" status="STATUS_UNDEFINED" >undefined</option>↵          </select>↵        </td>↵↵        <td valign="top" nowrap>↵          <label>progress</label><br>↵          <input type="text" name="progress" id="progress" size="7" class="formElements validated percentile" autocomplete="off" maxlength="255" value="" oldvalue="1" entrytype="PERCENTILE">↵        </td>↵      </tr>↵↵          </tr>↵          <tr>↵            <td colspan="4">↵              <label for="description">Description</label><br>↵              <textarea rows="3" cols="30" id="description" name="description" class="formElements" style="width:100%"></textarea>↵            </td>↵          </tr>↵        </table>↵↵    <h2>Assignments</h2>↵  <table  cellspacing="1" cellpadding="0" width="100%" id="assigsTable">↵    <tr>↵      <th style="width:100px;">name</th>↵      <th style="width:70px;">Role</th>↵      <th style="width:30px;">est.wklg.</th>↵      <th style="width:30px;" id="addAssig"><span class="teamworkIcon" style="cursor: pointer">+</span></th>↵    </tr>↵  </table>↵↵  <div style="text-align: right; padding-top: 20px">↵    <span id="saveButton" class="button first" onClick="$(this).trigger('saveFullEditor.gantt');">Save</span>↵  </div>↵↵  </div> -->↵  </div>↵↵↵↵<div class="__template__" type="ASSIGNMENT_ROW"><!--↵  <tr taskId="(#=obj.task.id#)" assId="(#=obj.assig.id#)" class="assigEditRow" >↵    <td ><select name="resourceId"  class="formElements" (#=obj.assig.id.indexOf("tmp_")==0?"":"disabled"#) ></select></td>↵    <td ><select type="select" name="roleId"  class="formElements"></select></td>↵    <td ><input type="text" name="effort" value="(#=getMillisInHoursMinutes(obj.assig.effort)#)" size="5" class="formElements"></td>↵    <td align="center"><span class="teamworkIcon delAssig del" style="cursor: pointer">d</span></td>↵  </tr>↵  --></div>↵↵↵↵<div class="__template__" type="RESOURCE_EDITOR">↵  <!-- <div class="resourceEditor" style="padding: 5px;">↵↵    <h2>Project team</h2>↵    <table  cellspacing="1" cellpadding="0" width="100%" id="resourcesTable">↵      <tr>↵        <th style="width:100px;">name</th>↵        <th style="width:30px;" id="addResource"><span class="teamworkIcon" style="cursor: pointer">+</span></th>↵      </tr>↵    </table>↵↵    <div style="text-align: right; padding-top: 20px"><button id="resSaveButton" class="button big">Save</button></div>↵  </div> -->↵  </div>↵↵↵↵<div class="__template__" type="RESOURCE_ROW"><!--↵  <tr resId="(#=obj.id#)" class="resRow" >↵    <td ><input type="text" name="name" value="(#=obj.name#)" style="width:100%;" class="formElements"></td>↵    <td align="center"><span class="teamworkIcon delRes del" style="cursor: pointer">d</span></td>↵  </tr>↵  --></div>↵↵↵";
    elems.each(function() {
      $(this).find(".__template__").each(function() {
          var tmpl = $(this);
          var type = tmpl.attr("type");

          //template may be inside <!-- ... --> or not in case of ajax loaded templates
          var found=false;
          var el=tmpl.get(0).firstChild;
          while (el && !found) {
            if (el.nodeType == 8) { // 8==comment
              var templateBody = el.nodeValue; // this is inside the comment
              found=true;
              break;
            }
            el=el.nextSibling;
          }
          if (!found){
            var templateBody = tmpl.html(); // this is the whole template
            
              if (type=="TASKROW"){
                var templateBody=' <tr taskId="(#=obj.id#)" class="taskEditRow (#=obj.isParent()?'+"'isParent':''#) (#=obj.collapsed?'collapsed':''#)"+'" level="(#=level#)">    <th class="gdfCell edit" align="right" style="cursor:pointer;"><span class="taskRowIndex">(#=obj.getRow()+1#)</span> <span class="teamworkIcon" style="font-size:8px;" >e</span></th>    <td class="gdfCell noClip" align="center"><div class="taskStatus cvcColorSquare" status="(#=obj.status#)"></div></td>    <td class="gdfCell"><input type="text" name="code" value="(#=obj.code?obj.code:'+"''#)"+'" placeholder="No."></td>    <td class="gdfCell indentCell" style="padding-left:(#=obj.level*10+18#)px;">      <div class="exp-controller" align="center"></div>      <input type="text" name="name" value="(#=obj.name#)" placeholder="Name">    </td>    <td class="gdfCell" align="center"><input type="checkbox" name="startIsMilestone"></td>    <td class="gdfCell"><input type="text" name="start"  value="" class="date"></td>    <td class="gdfCell" align="center"><input type="checkbox" name="endIsMilestone"></td>    <td class="gdfCell"><input type="text" name="end" value="" class="date"></td>    <td class="gdfCell"><input type="text" name="duration" autocomplete="off" value="(#=obj.duration#)"></td>    <td class="gdfCell"><input type="text" name="progress" class="validated" entrytype="PERCENTILE" autocomplete="off" value="(#=obj.progress?obj.progress:'+"''#)"+'" (#=obj.progressByWorklog?"readOnly":""#)></td>    <td class="gdfCell requireCanSeeDep"><input type="text" name="depends" autocomplete="off" value="(#=obj.depends#)" (#=obj.hasExternalDep?"readonly":""#)></td>    <td class="gdfCell taskAssigs">(#=obj.getAssigsString()#)</td>  </tr> ';
              }
              if (type=="TASKBAR"){
                var templateBody=' <div class="taskBox taskBoxDiv" taskId="(#=obj.id#)" >    <div class="layout (#=obj.hasExternalDep?'+"'extDep':''#)"+'">      <div class="taskStatus" status="(#=obj.status#)"></div>      <div class="taskProgress" style="width:(#=obj.progress>100?100:obj.progress#)%; background-color:(#=obj.progress>100?'+"'red':'rgb(153,255,51);'#);"+'"></div>      <div class="milestone (#=obj.startIsMilestone?'+"'active':''#)"+'" ></div>      <div class="taskLabel"></div>      <div class="milestone end (#=obj.endIsMilestone?'+"'active':''#)"+'" ></div>    </div>  </div> ';
              }
              if (type=="ASSIGNMENT_ROW"){
                var templateBody='  <tr taskId="(#=obj.task.id#)" assId="(#=obj.assig.id#)" class="assigEditRow" >    <td ><select name="resourceId"  class="formElements" (#=obj.assig.id.indexOf("tmp_")==0?"":"disabled"#) ></select></td>    <td ><select type="select" name="roleId"  class="formElements"></select></td>    <td ><input type="text" name="effort" value="(#=getMillisInHoursMinutes(obj.assig.effort)#)" size="5" class="formElements"></td>    <td align="center"><span class="teamworkIcon delAssig del" style="cursor: pointer">d</span></td>  </tr>  ';
              }
              if (type=="RESOURCE_ROW"){
                var templateBody='  <tr resId="(#=obj.id#)" class="resRow" >    <td ><input type="text" name="name" value="(#=obj.name#)" style="width:100%;" class="formElements"></td>    <td align="center"><span class="teamworkIcon delRes del" style="cursor: pointer">d</span></td>  </tr>  ';
              }
          }

          if (!templateBody.match(/##\w+##/)) { // is Resig' style? e.g. (#=id#) or (# ...some javascript code 'obj' is the alias for the object #)
            var strFunc =
                    "var p=[],print=function(){p.push.apply(p,arguments);};" +
                    "with(obj){p.push('" +
                    templateBody.replace(/[\r\t\n]/g, " ")
                            .replace(/'(?=[^#]*#\))/g, "\t")
                            .split("'").join("\\'")
                            .split("\t").join("'")
                            .replace(/\(#=(.+?)#\)/g, "',$1,'")
                            .split("(#").join("');")
                            .split("#)").join("p.push('")
                            + "');}return p.join('');";

          try {
            $.JST._templates[type] = new Function("obj", strFunc);
          } catch (e) {
            console.error("JST error: "+type, e,strFunc);
          }

          } else { //plain template   e.g. ##id##
          try {
            $.JST._templates[type] = templateBody;
          } catch (e) {
            console.error("JST error: "+type, e,templateBody);
          }
          }

          tmpl.remove();

      });
    });
  },

  createFromTemplate: function(jsonData, template, transformToPrintable) {
    var templates = $.JST._templates;

    var jsData=new Object();
    if (transformToPrintable){
      for (var prop in jsonData){
        var value = jsonData[prop];
        if (typeof(value) == "string")
          value = (value + "").replace(/\n/g, "<br>");
        jsData[prop]=value;
      }
    } else {
      jsData=jsonData;
    }

    function fillStripData(strip, data) {
      for (var prop in data) {
        var value = data[prop];

        strip = strip.replace(new RegExp("##" + prop + "##", "gi"), value);
      }
      // then clean the remaining ##xxx##
      strip = strip.replace(new RegExp("##\\w+##", "gi"), "");
      return strip;
    }

    var stripString = "";
    if (typeof(template) == "undefined") {
      alert("Template is required");
      stripString = "<div>Template is required</div>";

    } else if (typeof(templates[template]) == "function") { // resig template
      try {
        stripString = templates[template](jsData);// create a jquery object in memory
        if (template == "TASKEMPTYROW"){
          stripString =' <tr class="taskEditRow emptyRow" >     <th class="gdfCell" align="right"></th>     <td class="gdfCell noClip" align="center"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell"></td>     <td class="gdfCell requireCanSeeDep"></td>     <td class="gdfCell"></td>   </tr> '
        }
      } catch (e) {
        console.error("JST error: "+template,e.message);
        stripString = "<div> ERROR: "+template+"<br>" + e.message + "</div>";
      }

    } else {
      stripString = templates[template]; // recover strip template
      if (!stripString || stripString.trim() == "") {
        console.error("No template found for type '" + template + "'");
        return $("<div>");

      } else {
        stripString = fillStripData(stripString, jsData); //replace placeholders with data
      }
    }

    var ret = $(stripString);// create a jquery object in memory
    ret.attr("__template", template); // set __template attribute

    //decorate the strip
    var dec = $.JST._decorators[template];
    if (typeof (dec) == "function")
      dec(ret, jsData);

    return ret;
  },


  existsTemplate: function(template) {
    return $.JST._templates[template];
  },

  //decorate function is like function(domElement,jsonData){...}
  loadDecorator:function(template, decorator) {
    $.JST._decorators[template] = decorator;
  },

  getDecorator:function(template) {
    return $.JST._decorators[template];
  },

  decorateTemplate:function(element) {
    var dec = $.JST._decorators[element.attr("__template")];
    if (typeof (dec) == "function")
      dec(editor);
  },

  // asynchronous
  ajaxLoadAsynchTemplates: function(templateUrl, callback) {

    $.get(templateUrl, function(data) {

      var div = $("<div>");
      div.html(data);

      $.JST.loadTemplates(div);

      if (typeof(callback == "function"))
        callback();
    },"html");
  },

  ajaxLoadTemplates: function(templateUrl) {
    $.ajax({
      async:false,
      url: templateUrl,
      dataType: "html",
      success: function(data) {
        var div = $("<div>");
        div.html(data);
        $.JST.loadTemplates(div);
      }
    });

  }


};
