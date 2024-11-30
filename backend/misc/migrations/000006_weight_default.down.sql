UPDATE filters
SET weight = 1 
WHERE weight = 5 AND category = 'include';

UPDATE filters
SET weight = 2 
WHERE weight = 10 AND category = 'exclude';
