// The response of the "data" function is what's put in the database
function data() {
  return  {
    meta: {
      createdKey: 'created_at',
      updatedKey: 'updated_at'
    },
    data: [
      {
        item: 2,
        some_date: new Date(),
      },
      {
        item: 2,
        some_date: new Date(),
      },
    ]
  }
}
