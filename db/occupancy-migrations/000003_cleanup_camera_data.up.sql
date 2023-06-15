ALTER TABLE cameras DROP CONSTRAINT cameras_location_short_key;
ALTER TABLE cameras DROP CONSTRAINT cameras_location_long_key;

UPDATE cameras SET location_short = 'Science & Engineering' where location_short = 'SEL Main Entry C120';
UPDATE cameras SET location_short = 'Science & Engineering' where location_short = 'SEL Lower Level Stairwell C054';
UPDATE cameras SET location_short = 'Science & Engineering' where location_short = 'SEL Lower Level Stairwell C059';
UPDATE cameras SET location_short = 'Science & Engineering' where location_short = 'SEL Lower Level Stairwell C048';

UPDATE cameras SET location_short = 'Clemons' where location_short = '007 - 2nd Floor - West Exit';
UPDATE cameras SET location_short = 'Clemons' where location_short = '006 - 2nd Floor - North Exit';
UPDATE cameras SET location_short = 'Clemons' where location_short = '001A- 4th Floor - Main Camera 1';
UPDATE cameras SET location_short = 'Clemons' where location_short = '001D - 4th Floor - Main Camera 4';
UPDATE cameras SET location_short = 'Clemons' where location_short = '003 - 3rd Floor - Stairwell Exit';
UPDATE cameras SET location_short = 'Clemons' where location_short = '001B - 4th Fl - Main Camera 2';
UPDATE cameras SET location_short = 'Clemons' where location_short = '004 - 3rd Floor - Emergency Exit';
UPDATE cameras SET location_short = 'Clemons' where location_short = 'Axis-B8A44F10A1F7';
UPDATE cameras SET location_short = 'Clemons' where location_short = '002 - 3rd Floor - North Exit';
UPDATE cameras SET location_short = 'Clemons' where location_short = '001C - 4th Fl - Main Camera 3';
UPDATE cameras SET location_short = 'Clemons' where location_short = 'Axis-B8A44F07125E';
UPDATE cameras SET location_short = 'Clemons' where location_short = 'Axis-ACCC8EF00E06';
UPDATE cameras SET location_short = 'Clemons' where location_short = 'Axis-ACCC8EF02A0E';

UPDATE cameras SET location_short = 'Music' where location_short = 'OldCabell-B8A44F4F1939';
UPDATE cameras SET location_short = 'Fine Arts' where location_short = 'FiskeKimball-B8A44F4F195D';
