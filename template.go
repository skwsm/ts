package main

const TEMPLATE = `
<script src="//cdnjs.cloudflare.com/ajax/libs/polymer/0.5.2/polymer.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/handlebars.js/2.0.0/handlebars.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/lodash.js/2.4.1/lodash.min.js"></script>
<polymer-element name="togostanza-gene-attributes" attributes="gene_id">
<template>
<link rel="stylesheet" href="../assets/stanza.css">
<main id="content"></main>
</template>
<script>
  function Stanza(name, execute) {
    Polymer(name, {
      templates: {{ .TemplatesJson }},
      ready: function() {
        execute.bind(this)(this);
      },
      query: function(params, callback) {
        var self = this;
        var queryTemplate = Handlebars.compile(this.templates[params.template]);
        var query = queryTemplate(params.parameters);

        $.ajax({
          url: params.endpoint,
          data: {
            format: "json",
            query: query
          },
          success: function(data) {
            var rows = data.results.bindings.map(function(row) {
              var rowOut = {};
              _(row).forIn(function(value, key) {
                rowOut[key] = value.value;
              });
              return rowOut;
            });
            callback.bind(self)(rows);
          },
          error: function(xhr, status, err) {
            // TODO
          }
        });
      },
      render: function(params) {
        var htmlTemplate = Handlebars.compile(this.templates[params.template]);
        var htmlPartial = htmlTemplate(params.parameters);
        this.$.content.innerHTML = htmlPartial;
      }
    });
  }
{{ .IndexJs }}
</script>
</polymer-element>
`
