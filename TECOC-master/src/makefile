#
# This makefile builds TECOC on Windows using Microsoft compiler
# Before running it,  inspect the lines between here and the line containing
# all #'s to select the way you want TECO-C built.

# If you want to compile with built-in debugging support,  uncomment the
# following line.  When debugging is compiled in,  the control-P command
# can be used to turn on function tracing or get views of internal data
# structures (see the DbgDsp function in tecoc.c). Compiling in this support
# makes TECO-C larger and a little slower,  so you should only use this
# when you're debugging TECO-C.

#DEBG = -DDEBUGGING

# If you're debugging TECO-C,  use this option to compile in a consistency
# checking function called after each command executes.  The function checks
# the sanity of TECO-C's internal variables.  This obviously slows things
# down a little.

#CCHEK = -DCONSISTENCY_CHECKING

# If you want to use gcc,  uncomment the following lines,  which override the
# COMPILE.c macro so it doesn't stick in the -target option, which gcc
# doesn't understand.

# CC= cl -nologo

# Uncomment the following line for non-video teco
CFLAGS= -DWIN32 -DMSDOS

# Or, uncomment the following two lines for video teco
#CFLAGS= ${OSVERS} ${DEBG} ${CCHEK} -O  -DLINUX -DCURSES
#TERMOBJS = -lncurses


.c.obj:
	$(CC) $(CFLAGS) $(CPPFLAGS) -c $<

#############################################################################

OBJECTS=baksrc.obj bldstr.obj clenup.obj cmatch.obj docjr.obj doeves.obj doflag.obj \
	echoit.obj err.obj exea.obj exeats.obj exeb.obj exebar.obj exebsl.obj exec.obj \
	execcc.obj execln.obj execom.obj execrt.obj execst.obj execta.obj exectc.obj \
	exectd.obj execte.obj execti.obj exectl.obj exectn.obj execto.obj exectp.obj \
	exectq.obj exectr.obj exects.obj exectt.obj exectu.obj exectv.obj exectw.obj \
	exectx.obj execty.obj exectz.obj exed.obj exedgt.obj exedot.obj exedqu.obj \
	exee.obj exeequ.obj exeesc.obj exeexc.obj exeey.obj exef.obj exefb.obj exeg.obj \
	exegtr.obj exeh.obj exei.obj exeill.obj exej.obj exek.obj exel.obj exelbr.obj \
	exelst.obj exem.obj exen.obj exenul.obj exenyi.obj exeo.obj exeopr.obj exep.obj \
	exeprc.obj exepw.obj exeq.obj exeqes.obj exer.obj exerbr.obj exertp.obj exes.obj \
	exescl.obj exet.obj exeu.obj exeund.obj exeusc.obj exev.obj exew.obj exex.obj \
	exey.obj exez.obj findes.obj findqr.obj flowec.obj flowee.obj flowel.obj \
	getara.obj getnma.obj inccbp.obj init.obj insstr.obj isradx.obj ln2chr.obj \
	makdbf.obj makrom.obj popmac.obj pshmac.obj pushex.obj rdline.obj rdpage.obj \
	readcs.obj replac.obj search.obj singlp.obj skpcmd.obj srclop.obj sserch.obj \
	tabort.obj typbuf.obj typest.obj uminus.obj wrpage.obj zfrsrc.obj zwin32.obj

tecoc.exe: tecoc.obj $(OBJECTS)
	link -nologo tecoc.obj $(OBJECTS)

tecoc.obj: tecoc.c zport.h tecoc.h deferr.h dchars.h chmacs.h

clpars.h: genclp.exe clpars.tec
	genclp

genclp.exe: genclp.obj
	link /nologo /out:$@ genclp.obj

genclp.obj: genclp.c

clean:
	del *.obj
	del tecoc.exe tecoc.obj core genclp.obj genclp.exe clpars.h *.tmp *~ *.bak

tags:
	etags *.c *.h


#
#  Invoking this target (with "make depend") causes a new version of this
#  file to be created.  This current version will be saved in a file named
#  "makefile.bak".  In the new file, the lines following the special "DO NOT
#  DELETE" line will be replaced with a new set of dependency rules.
#

depend: makedep
	@echo '/^# DO NOT DELETE THIS LINE/+1,$$d' >makedep0
	@echo '$$r makedep5' >>makedep0
	@echo 'w' >>makedep0
	@cp makefile makefile.bak
	@ed - makefile < makedep0
	@rm makedep?
	@echo "new version of makefile generated, old version in makefile.bak"


# DO NOT DELETE THIS LINE -- make depend uses it
baksrc.obj: baksrc.c zport.h tecoc.h defext.h dchars.h chmacs.h deferr.h
bldstr.obj: bldstr.c zport.h tecoc.h defext.h deferr.h dchars.h chmacs.h
clenup.obj: clenup.c zport.h tecoc.h defext.h
cmatch.obj: cmatch.c zport.h tecoc.h defext.h dchars.h chmacs.h deferr.h
docjr.obj: docjr.c zport.h tecoc.h defext.h deferr.h
doeves.obj: doeves.c zport.h tecoc.h defext.h dchars.h
doflag.obj: doflag.c zport.h tecoc.h defext.h
echoit.obj: echoit.c zport.h tecoc.h defext.h dchars.h
err.obj: err.c zport.h tecoc.h defext.h deferr.h dchars.h
exea.obj: exea.c zport.h tecoc.h defext.h deferr.h
exeats.obj: exeats.c zport.h tecoc.h defext.h
exeb.obj: exeb.c zport.h tecoc.h defext.h
exebar.obj: exebar.c zport.h tecoc.h defext.h deferr.h
exebsl.obj: exebsl.c zport.h tecoc.h defext.h chmacs.h
exec.obj: exec.c zport.h tecoc.h defext.h
execcc.obj: execcc.c zport.h tecoc.h defext.h deferr.h
execln.obj: execln.c zport.h tecoc.h defext.h
execom.obj: execom.c zport.h tecoc.h defext.h deferr.h
execrt.obj: execrt.c zport.h tecoc.h defext.h chmacs.h deferr.h
execst.obj: execst.c zport.h tecoc.h defext.h deferr.h
execta.obj: execta.c zport.h tecoc.h defext.h dchars.h
exectc.obj: exectc.c zport.h tecoc.h defext.h dchars.h
exectd.obj: exectd.c zport.h tecoc.h defext.h
execte.obj: execte.c zport.h tecoc.h defext.h
execti.obj: execti.c zport.h tecoc.h defext.h dchars.h
exectl.obj: exectl.c zport.h tecoc.h defext.h dchars.h
exectn.obj: exectn.c zport.h tecoc.h defext.h
execto.obj: execto.c zport.h tecoc.h defext.h
exectp.obj: exectp.c zport.h tecoc.h defext.h
exectq.obj: exectq.c zport.h tecoc.h defext.h
exectr.obj: exectr.c zport.h tecoc.h defext.h deferr.h
exects.obj: exects.c zport.h tecoc.h defext.h
exectt.obj: exectt.c zport.h tecoc.h defext.h deferr.h
exectu.obj: exectu.c zport.h tecoc.h defext.h deferr.h dchars.h
exectv.obj: exectv.c zport.h tecoc.h defext.h deferr.h
exectw.obj: exectw.c zport.h tecoc.h defext.h deferr.h
exectx.obj: exectx.c zport.h tecoc.h defext.h
execty.obj: execty.c zport.h tecoc.h defext.h
exectz.obj: exectz.c zport.h tecoc.h defext.h
exed.obj: exed.c zport.h tecoc.h defext.h deferr.h
exedgt.obj: exedgt.c zport.h tecoc.h defext.h deferr.h chmacs.h
exedot.obj: exedot.c zport.h tecoc.h defext.h
exedqu.obj: exedqu.c zport.h tecoc.h defext.h deferr.h chmacs.h
exee.obj: exee.c zport.h tecoc.h defext.h chmacs.h dchars.h deferr.h
exeequ.obj: exeequ.c zport.h tecoc.h defext.h dchars.h deferr.h
exeesc.obj: exeesc.c zport.h tecoc.h defext.h dchars.h
exeexc.obj: exeexc.c zport.h tecoc.h defext.h
exeey.obj: exeey.c zport.h tecoc.h defext.h
exef.obj: exef.c zport.h tecoc.h defext.h chmacs.h deferr.h
exefb.obj: exefb.c zport.h tecoc.h defext.h deferr.h
exeg.obj: exeg.c zport.h tecoc.h defext.h deferr.h
exegtr.obj: exegtr.c zport.h tecoc.h defext.h deferr.h
exeh.obj: exeh.c zport.h tecoc.h defext.h
exei.obj: exei.c zport.h tecoc.h defext.h dchars.h deferr.h
exeill.obj: exeill.c zport.h tecoc.h defext.h deferr.h
exej.obj: exej.c zport.h tecoc.h defext.h
exek.obj: exek.c zport.h tecoc.h defext.h
exel.obj: exel.c zport.h tecoc.h defext.h
exelbr.obj: exelbr.c zport.h tecoc.h defext.h deferr.h
exelst.obj: exelst.c zport.h tecoc.h defext.h deferr.h
exem.obj: exem.c zport.h tecoc.h defext.h deferr.h
exen.obj: exen.c zport.h tecoc.h defext.h deferr.h
exenul.obj: exenul.c zport.h tecoc.h defext.h
exenyi.obj: exenyi.c zport.h tecoc.h defext.h deferr.h
exeo.obj: exeo.c zport.h tecoc.h defext.h dchars.h deferr.h
exeopr.obj: exeopr.c zport.h tecoc.h defext.h
exep.obj: exep.c zport.h tecoc.h defext.h deferr.h
exeprc.obj: exeprc.c zport.h tecoc.h defext.h deferr.h
exepw.obj: exepw.c zport.h tecoc.h defext.h
exeq.obj: exeq.c zport.h tecoc.h defext.h deferr.h
exeqes.obj: exeqes.c zport.h tecoc.h defext.h
exer.obj: exer.c zport.h tecoc.h defext.h
exerbr.obj: exerbr.c zport.h tecoc.h defext.h deferr.h
exertp.obj: exertp.c zport.h tecoc.h defext.h deferr.h
exes.obj: exes.c zport.h tecoc.h defext.h dchars.h deferr.h
exescl.obj: exescl.c zport.h tecoc.h defext.h deferr.h
exet.obj: exet.c zport.h tecoc.h defext.h dchars.h
exeu.obj: exeu.c zport.h tecoc.h defext.h deferr.h
exeund.obj: exeund.c zport.h tecoc.h defext.h deferr.h
exeusc.obj: exeusc.c zport.h tecoc.h defext.h dchars.h
exev.obj: exev.c zport.h tecoc.h defext.h
exew.obj: exew.c zport.h tecoc.h defext.h deferr.h
exex.obj: exex.c zport.h tecoc.h defext.h deferr.h
exey.obj: exey.c zport.h tecoc.h defext.h deferr.h
exez.obj: exez.c zport.h tecoc.h defext.h
findes.obj: findes.c zport.h tecoc.h defext.h deferr.h
findqr.obj: findqr.c zport.h tecoc.h defext.h deferr.h chmacs.h
flowec.obj: flowec.c zport.h tecoc.h defext.h deferr.h
flowee.obj: flowee.c zport.h tecoc.h defext.h deferr.h
flowel.obj: flowel.c zport.h tecoc.h defext.h deferr.h
getara.obj: getara.c zport.h tecoc.h defext.h deferr.h
getnma.obj: getnma.c zport.h tecoc.h defext.h deferr.h
inccbp.obj: inccbp.c zport.h tecoc.h defext.h deferr.h
init.obj: init.c zport.h tecoc.h defext.h deferr.h
insstr.obj: insstr.c zport.h tecoc.h defext.h deferr.h
isradx.obj: isradx.c zport.h tecoc.h defext.h chmacs.h
ln2chr.obj: ln2chr.c zport.h tecoc.h defext.h dchars.h chmacs.h
makdbf.obj: makdbf.c zport.h tecoc.h defext.h
makrom.obj: makrom.c zport.h tecoc.h defext.h deferr.h
popmac.obj: popmac.c zport.h tecoc.h defext.h
pshmac.obj: pshmac.c zport.h tecoc.h defext.h deferr.h
pushex.obj: pushex.c zport.h tecoc.h defext.h dchars.h deferr.h
rdline.obj: rdline.c zport.h tecoc.h defext.h deferr.h dchars.h
rdpage.obj: rdpage.c zport.h tecoc.h defext.h deferr.h
readcs.obj: readcs.c zport.h tecoc.h defext.h dchars.h chmacs.h deferr.h
readcs.obj: dscren.h
replac.obj: replac.c zport.h tecoc.h defext.h dchars.h
search.obj: search.c zport.h tecoc.h defext.h deferr.h dchars.h
singlp.obj: singlp.c zport.h tecoc.h defext.h dchars.h
skpcmd.obj: skpcmd.c zport.h tecoc.h defext.h chmacs.h dchars.h deferr.h
srclop.obj: srclop.c zport.h tecoc.h defext.h deferr.h
sserch.obj: sserch.c zport.h tecoc.h defext.h
tabort.obj: tabort.c zport.h tecoc.h
typbuf.obj: typbuf.c zport.h tecoc.h defext.h dchars.h chmacs.h
typest.obj: typest.c zport.h tecoc.h defext.h dchars.h
uminus.obj: uminus.c zport.h tecoc.h defext.h
wrpage.obj: wrpage.c zport.h tecoc.h defext.h dchars.h deferr.h
zfrsrc.obj: zfrsrc.c zport.h tecoc.h defext.h dchars.h chmacs.h deferr.h
zwin32.obj: zwin32.c tecoc.h clpars.h dchars.h deferr.h
zwin32.obj: defext.h dscren.h vrbmsg.h
