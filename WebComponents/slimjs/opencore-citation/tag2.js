		Slim.tag(
			'my-tag',
			`<div>Web Component version</div>
			<div>Citation:</div>
			<div>
			<span bind>{{Authors}}.</span>
			<span bind>{{Year}}.</span>
			<span bind>{{Dataset}}</span>
			<span style="font-style: italic;" bind>{{Title}}.</span>
			<span bind>{{Archive}}</span>
			<span bind>{{Version}}</span>
			<span bind>Retrieved from {{PID}}</span>
			</div>`,
			class MyTag extends Slim {
				// your code here
				onBeforeCreated() {
					var element = document.getElementById('test');
					var jsonld = element.innerHTML;

					// var promises = jsonld.promises;
					// var promise = promises.flatten(doc);
					// promise.then(function (flattened) { this.myMessage = flattened }, function (err) { this.myMessage = flattened });

					var obj = JSON.parse(jsonld);

					this.Authors = obj.publisher.name
					this.Year = ""
					this.Dataset = obj.name
					this.Title = obj.description
					this.Archive = ""
					this.Version = ""
					this.PID = obj.url
				}
			})
