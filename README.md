<img src="https://raw.githubusercontent.com/hershyz/sourcetools/master/img/logo.png"/>
<p>Cross platform, general CLI developer tools that can be used when working with large projects and huge codebases.</p>
<p>Sourcetools utilizes its own <strong>library</strong>, which can be found <a href="https://www.github.com/srctools/codebase">here</a>.</p>

<h4>Commands:</h4>
<pre>
  <i>compress (path)</i>                   Prints the code in a file without line breaks, useful for reducing the size of JavaScript files
  <i>copy (existing path, new path)</i>    Copies a file
  <i>definitions (path, term)</i>          Finds all definitions of a variable or function name
  <i>linecount (path)</i>                  Counts the number of lines in a file
  <i>nocomment (path)</i>                  Displays a file without comments
  <i>parameters (path, function)</i>       Displays the parameters of a certian function
  <i>read (path)</i>                       Reads a file
  <i>structure (path)</i>                  Shows what functions variables are initialzed in
  <i>varvals (path)</i>                    Shows all values of variables in a file
</pre>
<br>
<h4>Binaries:</h4>
<p>
  Windows binaries have already been compiled for all utilities, and can be found in <code>sourcetools/bin</code>.
  Any utility can be compiled into an executable for macOS or Linux using: <code>go build (filename).go codebase.go</code>
</p>
