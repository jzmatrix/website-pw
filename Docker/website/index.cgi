#!/usr/bin/perl -w
################################################################################
use CGI;
use Data::Dumper;
################################################################################
print "Content-type: text/html\n\n";
################################################################################
my $q = CGI->new;
#print "Q -- " .Dumper($q) . "\n";
################################################################################
my $len = $q->{'param'}{'len'}[0];
if (!$len){$len = 32;}
################################################################################
my $qty = $q->{'param'}{'qty'}[0];
if (!$qty) { $qty = 1;}
################################################################################
if ($q->{'param'}{'ssecret'})
{
	print 'J9p@qL~^LxTQ2xi[xR{Fv0V]tyx&5TqJ';
	exit;
}
################################################################################
$noSym = $q->{'param'}{'nosym'}[0];
$alphaOnly = $q->{'param'}{'alphaonly'}[0];
$numOnly = $q->{'param'}{'numonly'}[0];
$noForm = $q->{'param'}{'noform'}[0];
if (!$noSym) { $noSym = 0; }
if (!$alphaOnly) { $alphaOnly = 0; }
if (!$numOnly) { $numOnly = 0; }
if (!$noForm) { $noForm = 0; }
################################################################################
if (!$noForm)
{
	print "<!-- Length: $len -->\n";
	print "<!-- Quantity: $qty -->\n";
	print "<!-- NoSymbols : $noSym -->\n";
	print "<!-- AlphaOnly: $alphaOnly -->\n";
	print "<!-- NumOnly: $numOnly -->\n";
}
################################################################################
my $pwString;
################################################################################
for (my $x = 1;$x<=$qty;$x++)
{
	my $genCMD = "head /dev/urandom -n 50 | /usr/bin/strings --bytes 1";
	my $genPW = `$genCMD`; chomp $genPW;
	################
	my @lines = split("\n", $genPW);
	$newStr = "";
	foreach (@lines)
	{
		$newStr .= $_;
	}
	################
	my $pwStr;
	foreach (my $cnt = 0; $cnt <=length($newStr);$cnt++)
	{
		my $q = substr($newStr,$cnt,1);
		my $q2 = ord($q);
		#####
		my $okChar = 1;
		#####
		if ((!$noSym) && (!$alphaOnly) && (!$numOnly))
		{
			if ($q2 eq "32") { $okChar = 0; }	# space
			if ($q2 eq "96") { $okChar = 0; }	# `
			if ($q2 eq "34") { $okChar = 0; }	# "
			if ($q2 eq "39") { $okChar = 0; }	# '
			if ($q2 eq "9") { $okChar = 0; }	# tab
			if ($q2 eq "10") { $okChar = 0; }	# cr
			if ($q2 eq "124") { $okChar = 0; }	# |
			if ($q2 eq "37") { $okChar = 0; }	# %
			if ($q2 eq "40") { $okChar = 0; }	# (
			if ($q2 eq "41") { $okChar = 0; }	# )
			if ($q2 eq "47") { $okChar = 0; }	# /
			if ($q2 eq "58") { $okChar = 0; }	# :
			if ($q2 eq "60") { $okChar = 0; }	# <
			if ($q2 eq "62") { $okChar = 0; }	# >
		}
		elsif ($alphaOnly)
		{
			$okChar  = 0;
			if (($q2 >= 65) && ($q2 <= 90)) { $okChar = 1;}
			if (($q2 >= 97) && ($q2 <= 122)) { $okChar = 1;}
		}
		elsif ($numOnly)
		{
			$okChar  = 0;
			if (($q2 >= 48) && ($q2 <= 57)) { $okChar = 1;}
		}
		else
		{
			$okChar  = 0;
			if (($q2 >= 48) && ($q2 <= 57)) { $okChar = 1;}
			if (($q2 >= 65) && ($q2 <= 90)) { $okChar = 1;}
			if (($q2 >= 97) && ($q2 <= 122)) { $okChar = 1;}
		}
		#####
		if ($okChar)
		{
			$pwStr .= $q;
		}
		#print "<P>Q -- $q -- $q2 :: $pwStr\n";
	}
	################
	$newStr = substr($pwStr,0,$len);
	################
	if ($pwString) { $pwString .= "\n"; }
	$pwString .= $newStr;

}
my $len2 = $len + 10;
my $col2 = $qty + 5;
if (!$noForm)
{
	print "<form method='GET' action='/' name='MAIN'>\n";
	print "<table align='center' border='0'>\n";
	print "<tr><td align='center'><b>Params:</b></td><td><table border='0'>";
	print "<tr><td>Quantity: <td><td><input type='text' name='qty' size='5' value='$qty'></td></tr>\n";
	print "<tr><td>Length: <td><td><input type='text' name='len' size='5' value='$len'></td></tr>\n";
	print "<tr><td>No Symbols: </td><td><input type='checkbox' name='nosym' value='1' "; if ($noSym) { print " checked ";} print "</td></tr>\n";
	print "<tr><td>Alpha Only: </td><td><input type='checkbox' name='alphaonly' value='1' "; if ($alphaOnly) { print " checked ";} print "</td></tr>\n";
	print "<tr><td>Num Only: </td><td><input type='checkbox' name='numonly' value='1' "; if ($numOnly) { print " checked ";} print "</td></tr>\n";
	print "<tr><td colspan='2'><center><input type='submit' value='Generate New'></td></tr>\n";
	print "</table></td></tr>\n";
	print "<tr><td align='center'><b>Password List:</b></td><td><textarea cols=$len2 rows=$col2>$pwString</textarea></td></tr>\n";
	print "</table>\n";
	print "</form>\n";
}
else
{
	print $pwString;
}
