How to install a package:

	create a file inside src/
	define it as a package 'package pkgName'
	run 'go install pkgName'
	It will now be compiled inside the pkg folder.

Note however for a function to be used inside an import it must start with a capital letter. (This defines that it is exportable.)