class AddSortToCatalogAndLink < ActiveRecord::Migration[5.0]
  def change
    add_column :catalogs, :sort_by, :integer, null: false, default: 0
    add_column :links, :sort_by, :integer, null: false, default: 0
  end
end
