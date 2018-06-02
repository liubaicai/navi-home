module ApplicationHelper

  def isDebug
    return (ENV.fetch('RAILS_ENV') { 'development' })=='development'
  end

end
