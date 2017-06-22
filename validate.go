function validate() {
  var sheet = SpreadsheetApp.getActiveSheet();
  /*var links = sheet.getRange("A1:A10").getValues();*/
  var results = [];
  for (var i = 0; i < links.length; i++) {
    var code = 0;
    if (/^http/i.test(links[i][0])) {
      try {
        code = UrlFetchApp.fetch(links[i][0]).getResponseCode();
      }
      catch(e) { 
        print("Response Code not Available");
      }
    }
    results.push([code === 200]);
  }    
  /*sheet.getRange("B1:B10").setValues(results);*/
}