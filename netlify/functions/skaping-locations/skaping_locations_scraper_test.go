package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrapSkapingLocations_NominalCase_ListOfSkapingLocations(t *testing.T) {
	// GIVEN
	rawHtml := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			windows[1] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/pra-loup/molanes\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/pra-loup/molanes\">http://www.skaping.com/pra-loup/molanes</a>"
			});
			markers[1] = new google.maps.Marker({
				position: new google.maps.LatLng(44.36097000, 6.60417300),
				animation: google.maps.Animation.DROP,
				map: map,
				title:" PRA LOUP - CLOS DU SERRE (1820m)"
			});
			markers[1].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[1].open(map, markers[1]);
				openedWindow = windows[1];
			});
			bounds.extend(markers[1].getPosition());
		

			windows[2] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/ski-macedonia/bistra-mavrovo\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/ski-macedonia/bistra-mavrovo\">http://www.skaping.com/ski-macedonia/bistra-mavrovo</a>"
			});
			markers[2] = new google.maps.Marker({
				position: new google.maps.LatLng(41.64767700, 20.73542600),
				animation: google.maps.Animation.DROP,
				map: map,
				title:"#skimacedonia Bistra-Mavrovo"
			});
			markers[2].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[2].open(map, markers[2]);
				openedWindow = windows[2];
			});
			bounds.extend(markers[2].getPosition());
		</script>
	</head>
	</html>`
	rawHtmlNoMatch := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			console.log('Charly Lima');
		</script>
	</head>
	</html>
	`
	// WHEN
	skapingLocations := ScrapSkapingLocations(&rawHtml)
	skapingLocationsNoMatch := ScrapSkapingLocations(&rawHtmlNoMatch)
	// THEN
	expectedSkapingLocation1 := SkapingLocation{
		url:      "http://www.skaping.com/pra-loup/molanes",
		position: LatLng{44.36097000, 6.60417300},
		title:    " PRA LOUP - CLOS DU SERRE (1820m)",
	}
	expectedSkapingLocation2 := SkapingLocation{
		url:      "http://www.skaping.com/ski-macedonia/bistra-mavrovo",
		position: LatLng{41.64767700, 20.73542600},
		title:    "#skimacedonia Bistra-Mavrovo",
	}
	expectedSkapingLocations := []SkapingLocation{
		expectedSkapingLocation1,
		expectedSkapingLocation2,
	}
	assert.Equal(t, expectedSkapingLocations, skapingLocations)
	assert.Empty(t, skapingLocationsNoMatch)
}

func TestScrapSkapingRawDataLocations_NominalCase_ListOfRawDataLocations(t *testing.T) {
	// GIVEN
	rawHtml := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			windows[1] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/pra-loup/molanes\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/pra-loup/molanes\">http://www.skaping.com/pra-loup/molanes</a>"
			});
			markers[1] = new google.maps.Marker({
				position: new google.maps.LatLng(44.36097000, 6.60417300),
				animation: google.maps.Animation.DROP,
				map: map,
				title:" PRA LOUP - CLOS DU SERRE (1820m)"
			});
			markers[1].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[1].open(map, markers[1]);
				openedWindow = windows[1];
			});
			bounds.extend(markers[1].getPosition());
		

			windows[2] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/ski-macedonia/bistra-mavrovo\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/ski-macedonia/bistra-mavrovo\">http://www.skaping.com/ski-macedonia/bistra-mavrovo</a>"
			});
			markers[2] = new google.maps.Marker({
				position: new google.maps.LatLng(41.64767700, 20.73542600),
				animation: google.maps.Animation.DROP,
				map: map,
				title:"#skimacedonia Bistra-Mavrovo"
			});
			markers[2].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[2].open(map, markers[2]);
				openedWindow = windows[2];
			});
			bounds.extend(markers[2].getPosition());
		</script>
	</head>
	</html>`
	rawHtmlNoMatch := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			console.log('Charly Lima');
		</script>
	</head>
	</html>
	`
	// WHEN
	rawDataLocations := ScrapSkapingRawDataLocations(&rawHtml)
	rawDataLocationsNoMatch := ScrapSkapingRawDataLocations(&rawHtmlNoMatch)
	// THEN
	expectedRawDataLocation1 := `new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/pra-loup/molanes\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/pra-loup/molanes\">http://www.skaping.com/pra-loup/molanes</a>"
			});
			markers[1] = new google.maps.Marker({
				position: new google.maps.LatLng(44.36097000, 6.60417300),
				animation: google.maps.Animation.DROP,
				map: map,
				title:" PRA LOUP - CLOS DU SERRE (1820m)"
			});`
	expectedRawDataLocation2 := `new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/ski-macedonia/bistra-mavrovo\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/ski-macedonia/bistra-mavrovo\">http://www.skaping.com/ski-macedonia/bistra-mavrovo</a>"
			});
			markers[2] = new google.maps.Marker({
				position: new google.maps.LatLng(41.64767700, 20.73542600),
				animation: google.maps.Animation.DROP,
				map: map,
				title:"#skimacedonia Bistra-Mavrovo"
			});`
	expectedRawDataLocations := []string{
		expectedRawDataLocation1,
		expectedRawDataLocation2,
	}
	assert.Equal(t, expectedRawDataLocations, rawDataLocations)
	assert.Empty(t, rawDataLocationsNoMatch)
}

func TestExractSkapingLocationUrl_NominalCase_String(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	rawDataLocationNoMatch := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" ferh=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	locationUrl := ExractSkapingLocationUrl(&rawDataLocation)
	locationUrlNoMatch := ExractSkapingLocationUrl(&rawDataLocationNoMatch)
	// THEN
	expectedLocationUrl := "http://www.skaping.com/banffgondola"
	assert.Equal(t, expectedLocationUrl, locationUrl)
	assert.Empty(t, locationUrlNoMatch)
}

func TestExractSkapingLatLng_NominalCase_LatLng(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	rawDataLocationInvalidFloats := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(alpha, bravo),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	latLng := ExtractSkapingLocationLatLng(&rawDataLocation)
	latLngInvalidFloats := ExtractSkapingLocationLatLng(&rawDataLocationInvalidFloats)
	// THEN
	assert.Equal(t, LatLng{51.14460700, -115.57476600}, latLng)
	assert.Equal(t, LatLng{0, 0}, latLngInvalidFloats)
}

func TestExractSkapingLocationTitle_NominalCase_String(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	locationTitle := ExtractSkapingLocationTitle(&rawDataLocation)
	// THEN
	expectedLocationTitle := "Banff Gondola"
	assert.Equal(t, expectedLocationTitle, locationTitle)
}

func TestToSkapingLocation_NominalCase_SkapingLocation(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	skapingLocation := ToSkapingLocation(&rawDataLocation)
	// THEN
	expectedSkapingLocation := SkapingLocation{
		url:      "http://www.skaping.com/banffgondola",
		position: LatLng{51.14460700, -115.57476600},
		title:    "Banff Gondola",
	}
	assert.Equal(t, expectedSkapingLocation, skapingLocation)
}
