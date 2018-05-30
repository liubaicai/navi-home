class AddIconToLink < ActiveRecord::Migration[5.0]
  def change
    add_column :links, :icon, :string, default: ''
  end
end
