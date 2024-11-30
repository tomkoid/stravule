-- update the weight value to 5 if it is 1 and in include category
UPDATE filters
SET weight = 5
WHERE weight = 1 AND category = 'include';

-- update the weight value to 10 if it is 2 and in exclude category
UPDATE filters
SET weight = 10
WHERE weight = 2 AND category = 'exclude';
