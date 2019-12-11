'use strict';
const fs = require('fs');
const { Categories } = require('./CategoriesAndStyles.json');
const SimilarStyles = require('./similar-styles.json');

const categories = [];
const styles = [];

let catQueries = [];
let styleQueries = [];

const FILE_NAME = './styles/styles-db.sql';

let catId = 0;
let stId = 0;

const similar_styles = SimilarStyles.map(([a, b]) => `  (${a}, ${b})`);

Categories.forEach( ({ name, notes, type, _id: category_id, subcategory}) => {
  catId += 1;
  const cat = {
    id: catId,
    category_id,
    type,
    name: name.replace(/'/g, `''`),
    notes: notes.replace(/'/g, `''`),
  };
  categories.push(cat);
  catQueries.push(`  (${cat.id}, '${cat.category_id}', '${cat.type}', '${cat.name}', '${cat.notes}')`);


  subcategory.forEach(sub => {
    const {
      name, aroma, appearance, flavor, mouthfeel, impression, comments, history, ingredients, examples, varieties, tags, stats, comparison, entry_instructions, _id: style_id,
    } = sub;
    stId += 1;

    delete stats.ibu._flexible;
    delete stats.og._flexible;
    delete stats.fg._flexible;
    delete stats.srm._flexible;
    delete stats.abv._flexible;

    const s = {
      id: stId,
      category_id: cat.id,
      style_id,
      name,
      aroma: aroma ? aroma.replace(/'/g, `''`) : '',
      appearance: appearance ? appearance.replace(/'/g, `''`) : '',
      flavor: flavor ? flavor.replace(/'/g, `''`) : '',
      mouthfeel: mouthfeel ? mouthfeel.replace(/'/g, `''`) : '',
      impression: impression ? impression.replace(/'/g, `''`) : '',
      comments: comments ? comments.replace(/'/g, `''`) : '',
      history: history ? history.replace(/'/g, `''`) : '',
      ingredients: ingredients ? ingredients.replace(/'/g, `''`) : '',
      comparison: comparison ? comparison.replace(/'/g, `''`) : '',
      entry_instructions: entry_instructions ? entry_instructions.replace(/'/g, `''`) : '',
      varieties: varieties ? varieties.replace(/'/g, `''`).split(', ') : [],
      examples: examples ? examples.replace(/'/g, `''`).split(', ') : [],
      tags: tags ? tags.replace(/'/g, `''`).split(', ') : [],
      stats: JSON.stringify(stats),
      beer_exam: Number(cat.category_id) <= 26,
    };
    styles.push(s);

    s.tags.forEach( tag => {
      if (tag.length > 24) {
        console.log(style_id, 'tag:', tag);
      }
    });

    s.examples.forEach( example => {
      if (example.length > 64) {
        console.log(style_id, 'examples:', example);
      }
    });

    s.varieties.forEach( variety => {
      if (variety.length > 24) {
        console.log(style_id, 'varieties:', variety);
      }
    });


    styleQueries.push(`  (${s.id}, ${s.category_id}, '${s.style_id}', '${s.name}', '${s.aroma}', '${s.appearance}', '${s.flavor}', '${s.mouthfeel}', '${s.impression}', '${s.comments}', '${s.history}', '${s.ingredients}', '${s.comparison}', '${s.entry_instructions}', '{${s.examples.map(v => `"${v}"`).join(', ')}}', '{${s.varieties.map(v => `"${v}"`).join(', ')}}', '{${s.tags.map(v => `"${v}"`).join(', ')}}', '${s.stats}', '${s.beer_exam}')`);
  });
});


const categoryQueries = `INSERT INTO categories (id, category_id, type, name, notes) VALUES \n${catQueries.join(',\n')};\n\n`;
const subCatQueries = `INSERT INTO styles (id, category_id, style_id, name, aroma, appearance, flavor, mouthfeel, impression, comments, history, ingredients, comparison, entry_instructions, examples, varieties, tags, stats, beer_exam) VALUES \n${styleQueries.join(',\n')};\n\n`;
const similarStyleQueries = `INSERT INTO similar_styles (style_id, similar_id) VALUES \n${similar_styles.join(',\n')};\n\n`;


// fs.writeFileSync('./styles/categories-db.json', JSON.stringify(categories, null, 2));
// fs.writeFileSync('./styles/styles-db.json', JSON.stringify(styles, null, 2));
fs.writeFileSync(FILE_NAME, categoryQueries + subCatQueries + similarStyleQueries);


console.log(`${FILE_NAME} created successfully`);
