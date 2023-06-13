INSERT INTO public.abilities( id, ability) VALUES
(1, 'C') 
(2, 'C++') 
(3, 'C#') 
(4, 'Python') 
(5, 'HTML')  
(6, 'JavaScript');

INSERT INTO public.universities( id, university) VALUES
(1, 'Sakarya University') 
(2, 'Ataturk University') 
(3, 'Cumhuriyet University') 
(4, 'Trakya University') 
(5, 'Anadolu University')  
(6, 'Istanbul University');

INSERT INTO
    public.personals(
        id,
        name,
        surname,
        username,
        email,
        password,
        usertype
    )
VALUES (1, 'Abdullah', 'Gungor', 'abdullahgungor', 'abdullahgungor@gmail.com', '123456', 'Admin');

INSERT INTO public.personal_ability(ability_id, personal_id) VALUES
(1, 1) 
(2, 1) 
(3, 1) 
(4, 1) 
(5, 1)  
(6, 1);

INSERT INTO public.personal_university(university_id, personal_id) VALUES
(1, 1) 
(2, 1);

INSERT INTO
    public.experiences(
        id,
        company,
        "position",
        start_year,
        finish_year,
        personal_id
    )
VALUES 
(1, 'Ledzon', 'Sales Support Enginner', '2019-08-01' , '2020-02-14' , 1),
(2, 'EMFA', 'Sales Support Enginner', '2020-02-14' , '2021-09-30' , 1),
(3, 'Grup Arge', 'System Engineer', '2021-10-01', '2023-03-06', 1);