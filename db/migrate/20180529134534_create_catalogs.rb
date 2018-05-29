class CreateCatalogs < ActiveRecord::Migration[5.0]
  def change
    create_table :catalogs do |t|
      t.string :title,              default: ""
      t.belongs_to :user,           index: true

      t.timestamps
    end
  end
end
