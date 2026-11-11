--select avg(engagement) as engagement_avg from advertisements
  
--select count(id) as advertise_count from advertisements where engagement > 150

--select a.*, p.name from advertisements a join products p on a.product_id  = p.id  

--select max(engagement) as engagement_avg from advertisements

--select min(engagement) as engagement_avg from advertisements

--select p.*, a.title from products p left join advertisements a  on a.product_id  = null