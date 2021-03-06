class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable
  has_many :catalogs, -> { order 'sort_by asc' }, dependent: :destroy
  has_many :links, -> { order 'sort_by asc' }, dependent: :destroy
  validates :username, presence: true

end
