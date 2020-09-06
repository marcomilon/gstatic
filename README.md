# gStatic

A simple static website generator. The goal of the project is to create simple website generator. 
gStatic uses Go’s html/template and text/template libraries as the basis for the templating.

## Todo

- [x] Read all file on the template folder
- [x] Parse a yaml files and replace variables on the template
- [x] Write result to a public folder
- [x] Add support for layouts
- [x] Improve error handling
- [x] Add configuration
- [ ] Add documentation on how to do templates

## Common language

* sourceFolder is the folder that has all the html files.
* targetFolder is the folder where the website will be written.
* Template is an Html file.
* Data-source is a file that variables to be use inside the html files.

## How gStatic finds the data-source for an Html file

gStatic use one convention. 

> The name of the data-source must match the name of the html file.

Yaml files is the only data-source supported.

Example

| Template | Data-source | Result  |
| ------------- |-------------|-----|
| index.html | index.yaml | index.html will be rendered using variables on index.yaml |
| aboutus.html | aboutus.yaml | aboutus.html will be rendered using variables on aboutus.yaml |
| contactus.html | if no yaml file for contactus | contactus.html will be copied to Target as it is |

## Usage

> gstatic &lt;sourceFolder&gt; &lt;targetFolder&gt;

Use -h for help

> gstatic -h 