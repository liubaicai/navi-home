class Catalog < ApplicationRecord
  belongs_to :user
  has_many :links, -> { order 'sort_by asc' }, dependent: :destroy

end
