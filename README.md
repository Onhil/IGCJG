A web service to browse and upload .igc files


It's an RESTful API with 5 available calls:

- `GET /api`
	API status
- `GET /api/igc`
	Array of all stored tracks
- `POST /api/igc`
	Takes `{"url":"<url>"}` as JSON data and returns the assigned ID.
- `GET /api/igc/<id>`
	Returns track data (fields) for a valid `id`
- `GET /api/igc/<id>/<field>`
	Returns track `field` for a valid `id` and `field`
	
Demo of app
 - `https://igcjg.herokuapp.com/`
