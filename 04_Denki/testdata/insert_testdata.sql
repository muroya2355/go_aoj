DELETE FROM SUPER_VISOR;
INSERT INTO SUPER_VISOR VALUES('a', 'password1');
INSERT INTO SUPER_VISOR VALUES('A123456789', 'password2');

DELETE FROM CLASS;
INSERT INTO CLASS VALUES(1, '�①��');
INSERT INTO CLASS VALUES(2, '�e���r');
INSERT INTO CLASS VALUES(3, '����@');
INSERT INTO CLASS VALUES(4, '�G�A�R��');
INSERT INTO CLASS VALUES(5, '��C����@');

DELETE FROM MAKER;
INSERT INTO MAKER VALUES(1, '���A');
INSERT INTO MAKER VALUES(2, '���B');
INSERT INTO MAKER VALUES(3, '���C');
INSERT INTO MAKER VALUES(4, '���D');
INSERT INTO MAKER VALUES(5, '���E');

DELETE FROM GOODS;
INSERT INTO GOODS VALUES(1, '�T���v�����i', 1, 1, 'SAMPLE-GOODS-1', '�T���v���̏ڍ�', 100, 80.5, 1000, false, 'test1', current_timestamp, 1);
INSERT INTO GOODS VALUES(2, '�T���v�����O���i', 2, 2, 'SAMPLE-GOODS-2', '�T���v�����O�̏ڍ�', 1200, 1000, 20000, true, 'test2', current_timestamp, 1);
INSERT INTO GOODS VALUES(3, '���j�^���i', 3, 3, 'SAMPLE-GOODS-3', '���j�^�̏ڍ�', 10500, 9800, 300000, false, 'test3', current_timestamp, 1);
INSERT INTO GOODS VALUES(4, '���j�^�����O���i', 4, 4, 'SAMPLE-GOODS-4', '���j�^�����O�̏ڍ�', 50, 10.1, 100, true, 'test4', current_timestamp, 1);
