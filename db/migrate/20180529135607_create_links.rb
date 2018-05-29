class CreateLinks < ActiveRecord::Migration[5.0]
  def change
    create_table :links do |t|
      t.string :title,              default: ""
      t.belongs_to :user,           index: true
      t.belongs_to :catalog,        index: true

      t.timestamps
    end
  end
end
