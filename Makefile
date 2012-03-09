# par is made from par.6 
gopar : gopar.6 
	6l -o $@ $^ 
# we know that par.6 cannot be built without file 
gopar.6 : par2.a
# file.a is made from file.6 
par2.a : par2.6 
        # gopack never removes existing files from an archive 
        # so better remove the whole archive beforehand. 
		[ ! -e $@ ] || rm $@ 
		gopack grc $@ $^ 
# a generic rule for compiling any .go file into .6 
%.6 : %.go 
	6g -o $@ $<	
clean:
	rm gopar par2.a par2.6 gopar.6
