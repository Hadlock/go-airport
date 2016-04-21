/* appstatus.js
 * Copyright 2016 UpGuard Inc.
 * Zakk Acreman <zakk.acreman@upguard.com>
 * All rights reserved
 */

function populateDockerStats(data) {
		$("#docker-failure").remove();

		var table = $('#docker-table');		

		function newTd (value) {
				return $(document.createElement('td')).append(value);
		}
		
		data.forEach(function (v, i, data) {
				var id = v["Id"].slice(0,8),
						name = v["Names"][0],
						image = v["Image"],
						status = v["Status"].split(" ")[0],
						time = v["Status"].split(" ").slice(1).join(" ");
				
				table.append($(document.createElement('tr'))
										 .append(newTd(id))
										 .append(newTd(name))
										 .append(newTd(image))
										 .append(newTd(status))
										 .append(newTd(time)))
		});
		table.removeClass("hidden");
}

function populateFleetUnits(data) {
		$("#fleet-failure").remove();
}
